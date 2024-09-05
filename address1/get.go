package address1

import (
	"context"
	"github.com/advanced-go/stdlib/core"
	"net/http"
	"net/url"
)

func get[E core.ErrorHandler](ctx context.Context, h http.Header, values url.Values) (entries []Entry, h2 http.Header, status *core.Status) {
	var e E

	if values == nil {
		return nil, h2, core.StatusNotFound()
	}

	// TODO : get entries
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
		return nil, h2, status
	}
	if values != nil && len(values) > 0 {
		entries = filter(entries, values)
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

func filter(entries []Entry, values url.Values) (result []Entry) {
	customer := values.Get("customer")
	for _, e := range entries {
		if customer != "" && customer != e.CustomerId {
			continue
		}
		result = append(result, e)
	}
	return result
}
