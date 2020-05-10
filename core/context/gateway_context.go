package context

import (
	"context"
	"net/http"
	"net/url"
)

type GatewayContext struct {
	ResponseWriter http.ResponseWriter
	Request   *http.Request
	Context   context.Context
	Body      []byte
	Response      *http.Response
	RemoteURL *url.URL
}

func NewContext(w http.ResponseWriter,r *http.Request) *GatewayContext {
	return &GatewayContext{
		w,
		r,
		r.Context(),
		nil,
		nil,
		nil,
	}
}
