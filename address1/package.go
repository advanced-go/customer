package address1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/customer/address1"

	StorageHost  = "www.documents.com"
	StoragePath  = "storage/address"
	StorageRoute = "customer-address"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Get - address1 resource GET
func Get(r *http.Request, _ string) (entries []Entry, h2 http.Header, status *core.Status) {
	if r == nil {
		return entries, h2, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
	}
	return get[core.Output](r.Context(), r.Header, r.URL.Query())
}

// Put - address1 PUT, with optional content override
func Put(r *http.Request, _ string, body []Entry) (http.Header, *core.Status) {
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
