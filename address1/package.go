package address1

import (
	"errors"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	json2 "github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/customer/address1"

	StorageHost  = "www.documents.com"
	StoragePath  = "storage/address"
	StorageRoute = "customer-address"
	addrPath     = "address/entry"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Get - address1 resource GET
func Get[E core.ErrorHandler](r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	var e E

	h2 := httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText)
	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
		e.Handle(status)
		return nil, h2, status
	}
	switch path {
	case addrPath:
		t, status := get[E](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := json2.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson), status1
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not percentile or status code"))
		e.Handle(status)
		return nil, h2, status
	}
}

// Put - address1 PUT, with optional content override
func Put[E core.ErrorHandler](r *http.Request, _ string, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if body == nil {
		content, status := json2.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e E
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}
