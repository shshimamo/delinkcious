package user_client

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

func NewClient(baseURL string) (om.UserManager, error) {
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "http://" + baseURL
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	registerEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/register"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse).Endpoint()

	loginEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/login"),
		encodeHTTPGenericRequest,
		decodeLoginResponse).Endpoint()

	logoutEndpoint := httptransport.NewClient(
		"POST",
		copyURL(u, "/logout"),
		encodeHTTPGenericRequest,
		decodeSimpleResponse).Endpoint()

	return EndpointSet{
		RegisterEndpoint: registerEndpoint,
		LoginEndpoint:    loginEndpoint,
		LogoutEndpoint:   logoutEndpoint,
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
	return nil
}
