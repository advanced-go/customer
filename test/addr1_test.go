package test

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/test"
	"github.com/advanced-go/customer/address1"
	http2 "github.com/advanced-go/customer/http"
	"github.com/advanced-go/customer/testrsc"
	"net/http"
	"reflect"
	"testing"
)

func TestExchange1(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		{name: "get-error-header", req: test.NewRequestTest(testrsc.Addr1GetReqErrHeader, t), resp: test.NewResponseTest(testrsc.Addr1GetResp, t), status: core.StatusOK()},
		{name: "get-error-content", req: test.NewRequestTest(testrsc.Addr1GetReq, t), resp: test.NewResponseTest(testrsc.Addr1GetRespErrContent, t), status: core.StatusOK()},
		{name: "get-entry", req: test.NewRequestTest(testrsc.Addr1GetReq, t), resp: test.NewResponseTest(testrsc.Addr1GetResp, t), status: core.StatusOK()},
	}
	for _, tt := range tests {
		ok := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, error : %v", status.Code, tt.status.Code, status.Err)
				ok = false
			}
			if ok && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				ok = false
			}
			var gotT []address1.Entry
			var wantT []address1.Entry
			if ok {
				gotT, wantT, ok = test.Deserialize[test.Output, []address1.Entry](resp.Body, tt.resp.Body, t)
			}
			if ok {
				if !reflect.DeepEqual(gotT, wantT) {
					t.Errorf("Exchange() got = %v, want %v", gotT, wantT)
				}
			}
		})
	}
}
