package address1

import (
	"fmt"
	"github.com/advanced-go/customer/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/url"
)

func ExampleGet() {
	values := make(url.Values)
	values.Add(CustomerKey, "C001")
	path := uri.BuildPath(UpstreamPath, values)
	h := uri.AddContentLocation(nil, path, testrsc.Addr1GetRespURL)

	entries, _, status := get[core.Output](nil, h, values)
	fmt.Printf("test: get() -> [status:%v] [path:%v] [entries:%v]\n", status, path, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:1]

}
