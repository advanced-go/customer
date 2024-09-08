package address1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

const (
	PkgPath      = "github/advanced-go/customer/address1"
	UpstreamPath = "storage/address"
	CustomerKey  = "customer"
	Route        = "cust-address"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Url - egress URLs
func Url(host, path string, values url.Values, h http.Header) string {
	return resolver.Url(host, path, values, h)
}

// Get - timeseries2 resource GET
func Get(ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	return get[core.Output](ctx, h, values)
}

// Put - timeseries2 PUT, with optional content override
func Put(r *http.Request, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e core.Log
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}
