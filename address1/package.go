package address1

import (
	"context"
	"errors"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/jsonx"
	"github.com/advanced-go/common/uri"
	"net/http"
	"net/url"
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
func Get(r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	if r == nil {
		status := core.NewStatusError(core.StatusInvalidArgument, errors.New("error: http.Request is"))
		return nil, nil, status
	}
	if r.Header.Get(core.XFrom) == "" {
		return httpGet[core.Log](r, path)
	}
	return httpGet[core.Output](r, path)
}

func httpGet[E core.ErrorHandler](r *http.Request, path string) ([]byte, http.Header, *core.Status) {
	var e E

	switch path {
	case addrPath:
		t, h2, status := get[E](r.Context(), r.Header, r.URL.Query())
		if !status.OK() {
			return nil, h2, status
		}
		buf, status1 := jsonx.Marshal(t)
		if !status1.OK() {
			e.Handle(status1)
			return nil, h2, status1
		}
		return buf, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeJson), status1
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New("error: resource is not percentile or status code"))
		return nil, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText), status
	}
}

// Put - log2 PUT, with optional content override
func Put(r *http.Request, path string, body []Entry) (http.Header, *core.Status) {
	if r == nil {
		return nil, core.NewStatusError(core.StatusInvalidArgument, errors.New("error: request is nil"))
	}
	if r.Header.Get(core.XFrom) == "" {
		return httpPut[core.Log](r, path, body)
	}
	return httpPut[core.Output](r, path, body)
}

func httpPut[E core.ErrorHandler](r *http.Request, _ string, body []Entry) (http.Header, *core.Status) {
	if body == nil {
		content, status := jsonx.New[[]Entry](r.Body, r.Header)
		if !status.OK() {
			var e E
			e.Handle(status.WithRequestId(r.Header))
			return nil, status
		}
		body = content
	}
	return put[core.Log](r.Context(), core.AddRequestId(r.Header), body)
}

func AddressQuery(ctx context.Context, h http.Header, values url.Values) ([]Entry, *core.Status) {
	e, _, status := get[core.Log](ctx, h, values)
	return e, status
}
