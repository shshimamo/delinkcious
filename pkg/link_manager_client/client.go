package link_manager_client

import (
	"bytes"
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	om "github.com/shshimamo/delinkcious/pkg/object_model"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func NewClient(baseURL string) (om.LinkManager, error) {
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	getLinksEndpoint := httptransport.NewClient(
		"GET",
		copyURL(u, "/links"),
		encodeGetLinksRequest,
		decodeGetLinksResponse).Endpoint()

	addLinkEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/links"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse).Endpoint()

	updateLinkEndpoint := httptransport.NewClient(
		"PUT",
		copyURL(u, "/links"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse).Endpoint()

	deleteLinkEndpoint := httptransport.NewClient(
		"DELETE",
		copyURL(u, "/links"),
		encodeDeleteLinkRequest,
		decodeSimpleResponse).Endpoint()

	return EndpointSet{
		GetLinksEndpoint:   getLinksEndpoint,
		AddLinkEndpoint:    addLinkEndpoint,
		UpdateLinkEndpoint: updateLinkEndpoint,
		DeleteLinkEndpoint: deleteLinkEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base // メモリ値をコピーする
	next.Path = path
	return &next
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = io.NopCloser(&buf)
	return nil
}

func encodeGetLinksRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(om.GetLinksRequest)
	urlRegex := url.QueryEscape(r.UrlRegex)
	titleRegex := url.QueryEscape(r.TitleRegex)
	descriptionRegex := url.QueryEscape(r.DescriptionRegex)
	username := url.QueryEscape(r.Username)
	tag := url.QueryEscape(r.Tag)
	startToken := url.QueryEscape(r.StartToken)

	q := req.URL.Query()
	q.Add("url", urlRegex)
	q.Add("title", titleRegex)
	q.Add("description", descriptionRegex)
	q.Add("username", username)
	q.Add("tag", tag)
	q.Add("start", startToken)
	req.URL.RawQuery = q.Encode()
	return encodeHTTPGenericRequest(ctx, req, request)
}

func encodeDeleteLinkRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(*deleteLinkRequest)
	q := req.URL.Query()
	q.Add("username", r.Username)
	q.Add("url", r.Url)
	req.URL.RawQuery = q.Encode()
	return encodeHTTPGenericRequest(ctx, req, request)
}
