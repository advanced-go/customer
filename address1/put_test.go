package address1

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/jsonx"
	"github.com/advanced-go/common/uri"
	"github.com/advanced-go/customer/testrsc"
)

func ExamplePut() {
	entries, _ := jsonx.New[[]Entry](testrsc.Addr1Entry, nil)
	path := uri.BuildPath("", StoragePath, nil)
	h := uri.AddResolverEntry(nil, path, testrsc.Addr1PutRespFailure)

	_, status := put[core.Output](nil, h, entries)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:Timeout]

}
