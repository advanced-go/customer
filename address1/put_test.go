package address1

import (
	"fmt"
	"github.com/advanced-go/customer/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/uri"
)

func ExamplePut() {
	entries, _ := json.New[[]Entry](testrsc.Addr1Entry, nil)
	path := uri.BuildPath("", StoragePath, nil)
	h := uri.AddResolverContentLocation(nil, path, testrsc.Addr1PutRespFailure)

	_, status := put[core.Output](nil, h, entries)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:Timeout]

}
