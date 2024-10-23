package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/customer/address1"
	"github.com/advanced-go/customer/module"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/customer/http"
	ver1    = "v1"
	//Ver2            = "v2"
	address      = "address"
	AddressRoute = "cust-address"

	healthLivenessPath  = "health/liveness"
	healthReadinessPath = "health/readiness"
	versionPath         = "version"
	authorityPath       = "authority"
	//AuthorityRootPath   = "/authority"

)

var (
	authorityResponse = NewAuthorityResponse(module.Authority)
)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	p, status := httpx.ValidateURL(r.URL, module.Authority)
	if !status.OK() {
		resp, _ := httpx.NewResponse(status.HttpCode(), h2, status.Err)
		return resp, status
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case address:
		resp, status1 := addressExchange(r, p)
		resp.Header.Add(core.XRoute, address1.StorageRoute)
		return resp, status1
	case versionPath:
		return NewVersionResponse(module.Version), core.StatusOK()
	case authorityPath:
		return authorityResponse, core.StatusOK()
	case healthReadinessPath, healthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status = core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, testresource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}
