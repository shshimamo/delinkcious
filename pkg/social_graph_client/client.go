package social_graph_client

import (
	"bytes"
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/shshimamo/delinkcious/pkg/auth_util"
	om "github.com/shshimamo/delinkcious/pkg/object_model"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const SERVICE_NAME = "social-graph-manager"

func NewClient(baseURL string) (om.SocialGraphManager, error) {
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	followEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/follow"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse).Endpoint()

	unfollowEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/unfollow"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse).Endpoint()

	getFollowingEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/following"),
		encodeGetByUsernameRequest,
		decodeGetFollowingResponse).Endpoint()

	getFollowersEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/followers"),
		encodeGetByUsernameRequest,
		decodeGetFollowersResponse).Endpoint()

	return EndpointSet{
		FollowEndpoint:       followEndpoint,
		UnfollowEndpoint:     unfollowEndpoint,
		GetFollowingEndpoint: getFollowingEndpoint,
		GetFollowersEndpoint: getFollowersEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = io.NopCloser(&buf)

	if os.Getenv("DELINKCIOUS_MUTUAL_AUTH") != "false" {
		token := auth_util.GetToken(SERVICE_NAME)
		r.Header["Delinkcious-Caller-Token"] = []string{token}
	}

	return nil
}

func encodeGetByUsernameRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getByUserNameRequest)
	username := url.PathEscape(r.Username)
	req.URL.Path += "/" + username
	return encodeHTTPGenericRequest(ctx, req, request)
}
