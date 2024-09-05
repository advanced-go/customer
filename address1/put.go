package address1

import (
	"context"
	"errors"
	"github.com/advanced-go/stdlib/core"
	"net/http"
)

// put - function to Put a set of entries into a datastore
func put[E core.ErrorHandler](ctx context.Context, h http.Header, body []Entry) (h2 http.Header, status *core.Status) {
	var e E

	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	// TODO : put entries
	if !status.OK() {
		e.Handle(status, core.RequestId(h))
	}
	return
}
