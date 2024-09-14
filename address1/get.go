package address1

import (
	"context"
	"github.com/advanced-go/customer/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
	"net/url"
)

const (
	customerKey = "customer"
	stateKey    = "state"
)

func testOverride(h http.Header) http.Header {
	if h != nil && h.Get(uri.XContentLocationResolver) != "" {
		return h
	}
	return httpx.SetHeader(h, uri.XContentLocationResolver, testrsc.Addr1GetRespTest)
}

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(h2, httpx.ContentType, httpx.ContentTypeText)
	if values == nil {
		return nil, h2, core.StatusNotFound()
	}
	// Test only
	h = testOverride(h)

	u := resolver.Url(StorageHost, "", StoragePath, values, h)
	req, err := http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
	if err != nil {
		return nil, h2, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	resp, status1 := httpx.Exchange(req)
	if !status1.OK() {
		e.Handle(status1.WithRequestId(h))
		return nil, h2, status1
	}
	entries, status = json.New[[]Entry](resp, h)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
		return nil, h2, status
	}
	if len(values) > 0 {
		entries = filter(entries, values)
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	} else {
		h2 = httpx.SetHeader(h2, httpx.ContentType, httpx.ContentTypeJson)
	}
	return
}

func filter(entries []Entry, values url.Values) (result []Entry) {
	customer := values.Get(customerKey)
	state := values.Get(stateKey)
	for _, e := range entries {
		if customer == "*" {
			result = append(result, e)
			continue
		}
		if customer != "" && customer != e.CustomerId {
			continue
		}
		if state != "" && state != e.State {
			continue
		}
		result = append(result, e)
	}
	return result
}
