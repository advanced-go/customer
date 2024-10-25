package address1

import (
	"context"
	"errors"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
)

// put - function to Put a set of entries into a datastore
func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) (h2 http.Header, status *core.Status) {
	var e E

	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	u := resolver.Url(StorageHost, "", StoragePath, nil, h)
	req, err := http.NewRequestWithContext(core.NewContext(ctx), http.MethodGet, u, nil)
	if err != nil {
		return h2, core.NewStatusError(core.StatusInvalidArgument, err)
	}
	_, status = httpx.Exchange(req)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
	}
	return
}
