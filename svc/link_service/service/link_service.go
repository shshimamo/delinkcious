package service

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shshimamo/delinkcious/pkg/db_util"
	"github.com/shshimamo/delinkcious/pkg/link_manager_events"
	"github.com/shshimamo/delinkcious/pkg/log"
	om "github.com/shshimamo/delinkcious/pkg/object_model"
	"github.com/uber/jaeger-client-go"
	"io"
	"net/http"
	"os"
	"strconv"

	lm "github.com/shshimamo/delinkcious/pkg/link_manager"
	jeagerconfig "github.com/uber/jaeger-client-go/config"
)

type EventSink struct {
}

type linkManagerMiddleware func(om.LinkManager) om.LinkManager

func (s *EventSink) OnLinkAdded(username string, link *om.Link) {
}

func (s *EventSink) OnLinkUpdated(username string, link *om.Link) {
}

func (s *EventSink) OnLinkDeleted(username string, url string) {
}

func createTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jeagerconfig.Configuration{
		ServiceName: service,
		Sampler: &jeagerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jeagerconfig.ReporterConfig{
			LogSpans: true,
		},
	}
	logger := jeagerconfig.Logger(jaeger.StdLogger)
	tracer, closer, err := cfg.NewTracer(logger)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot create tracer: %v\n", err))
	}
	return tracer, closer
}

func Run() {
	dbHost, dbPort, err := db_util.GetDbEndpoint("link")
	if err != nil {
		log.Fatal(err)
	}
	store, err := lm.NewDbLinkStore(dbHost, dbPort, "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}

	sgHost := os.Getenv("SOCIAL_GRAPH_MANAGER_SERVICE_HOST")
	if sgHost == "" {
		sgHost = "localhost"
	}

	sgPort := os.Getenv("SOCIAL_GRAPH_MANAGER_SERVICE_PORT")
	if sgPort == "" {
		sgPort = "9090"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	maxLinksPerUserStr := os.Getenv("MAX_LINKS_PER_USER")
	if maxLinksPerUserStr == "" {
		maxLinksPerUserStr = "10"
	}

	maxLinksPerUser, err := strconv.ParseInt(maxLinksPerUserStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	socialGraphClient, err := sgm.NewClient(fmt.Sprintf("%s:%s", sgHost, sgPort))
	if err != nil {
		log.Fatal(err)
	}

	natsHostname := os.Getenv("NATS_CLUSTER_SERVICE_HOST")
	natsPort := os.Getenv("NATS_CLUSTER_SERVICE_PORT")

	natsUrl := ""
	var eventSink om.LinkManagerEvents
	if natsHostname != "" {
		natsUrl = natsHostname + ":" + natsPort
		eventSink, err = link_manager_events.NewEventSender(natsUrl)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		eventSink = &EventSink{}
	}

	// Create a logger
	logger := log.NewLogger("link manager")

	// Create a tracer
	tracer, closer := createTracer("link-manager")
	defer closer.Close()

	// Create the service implementation
	svc, err := lm.NewLinkManager(store, socialGraphClient, natsUrl, eventSink, maxLinksPerUser)
	if err != nil {
		log.Fatal(err)
	}

	svc = newLoggingMiddleware(logger)(svc)
	svc = newMetricsMiddleware()(svc)
	svc = newTracingMiddleware(tracer)(svc)

	getLinksHandler := httptransport.NewServer(
		makeGetLinksEndpoint(svc),
		decodeGetLinksRequest,
		encodeResponse,
	)

	addLinkHandler := httptransport.NewServer(
		makeAddLinkEndpoint(svc),
		decodeAddLinkRequest,
		encodeResponse,
	)

	updateLinkHandler := httptransport.NewServer(
		makeUpdateLinkEndpoint(svc),
		decodeUpdateLinkRequest,
		encodeResponse,
	)

	deleteLinkHandler := httptransport.NewServer(
		makeDeleteLinkEndpoint(svc),
		decodeDeleteLinkRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/links").Handler(getLinksHandler)
	r.Methods("POST").Path("/links").Handler(addLinkHandler)
	r.Methods("PUT").Path("/links").Handler(updateLinkHandler)
	r.Methods("DELETE").Path("/links").Handler(deleteLinkHandler)
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	logger.Log("msg", "*** listening on ***", "port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
