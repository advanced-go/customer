package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/uri"
	"github.com/advanced-go/customer/address1"
	"github.com/advanced-go/customer/module"
	"net/http"
)

func addressExchange(r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, status := httpx.ValidateURL(r.URL, module.Authority)
		if !status.OK() {
			return httpx.NewResponse(status.HttpCode(), h2, status.Err)
		}
		p = p1
	}

	switch r.Method {
	case http.MethodGet:
		return addressGet(r, p)
	case http.MethodPut:
		return addressPut(r, p)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}

func addressGet(r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var entries any
	var h2 http.Header

	switch p.Version {
	case ver1, "":
		entries, h2, status = address1.Get(r, p.Path)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if !status.OK() {
		resp, _ = httpx.NewResponse(status.HttpCode(), h2, status.Err)
		return
	}
	return httpx.NewResponse(status.HttpCode(), h2, entries)
}

func addressPut(r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch p.Version {
	case ver1, "":
		h2, status = address1.Put(r, p.Path, nil)
	//case module.Ver2:
	//	h2, status = address2.Put(r, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	return httpx.NewResponse(status.HttpCode(), h2, status.Err)
}
