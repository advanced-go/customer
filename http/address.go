package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/customer/address1"
	"github.com/advanced-go/customer/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

func addressExchange[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
		}
		p = p1
	}

	switch r.Method {
	case http.MethodGet:
		return addressGet[E](r, p)
	case http.MethodPut:
		return addressPut[E](r, p)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
	}
}

func addressGet[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var entries any
	var h2 http.Header

	switch p.Version {
	case module.Ver1, "":
		entries, h2, status = address1.Get(r, p.Path)
	//case module.Ver2:
	//entries, h2, status = timeseries2.Get(ctx, h, url.Query())
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if h2 == nil {
		h2 = make(http.Header)
	}
	if !status.OK() {
		h2.Add(httpx.ContentType, httpx.ContentTypeText)
		resp, _ = httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
		return
	}
	h2.Add(httpx.ContentType, httpx.ContentTypeJson)
	return httpx.NewResponse[E](status.HttpCode(), h2, entries)

}

func addressPut[E core.ErrorHandler](r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch p.Version {
	case module.Ver1, "":
		h2, status = address1.Put(r, p.Path, nil)
	//case module.Ver2:
	//	h2, status = address2.Put(r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if h2 == nil {
		h2 = make(http.Header)
	}
	h2.Add(httpx.ContentType, httpx.ContentTypeText)
	return httpx.NewResponse[E](status.HttpCode(), h2, status.Err)
}
