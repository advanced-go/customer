package address1

import (
	"fmt"
	"github.com/advanced-go/customer/testrsc"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/uri"
	"net/url"
)

func ExampleGet_Customer() {
	values := make(url.Values)
	values.Add(customerKey, "C001")
	path := uri.BuildPath("", StoragePath, values)
	h := uri.AddResolverContentLocation(nil, path, testrsc.Addr1GetRespTest)

	entries, status := get[core.Output](nil, h, values)
	fmt.Printf("test: get() -> [status:%v] [path:%v] [entries:%v]\n", status, path, len(entries))

	//Output:
	//test: get() -> [status:OK] [path:storage/address?customer=C001] [entries:1]

}

func ExampleGet_Customer_All() {
	values := make(url.Values)
	values.Add(customerKey, "*")
	path := uri.BuildPath("", StoragePath, values)
	h := uri.AddResolverContentLocation(nil, path, testrsc.Addr1GetRespTest)

	entries, status := get[core.Output](nil, h, values)
	fmt.Printf("test: get() -> [status:%v] [path:%v] [entries:%v]\n", status, path, len(entries))

	//Output:
	//test: get() -> [status:OK] [path:storage/address?customer=*] [entries:4]

}

func ExampleGet_State() {
	values := make(url.Values)
	values.Add(stateKey, "IA")
	path := uri.BuildPath("", StoragePath, values)
	h := uri.AddResolverContentLocation(nil, path, testrsc.Addr1GetRespTest)

	entries, status := get[core.Output](nil, h, values)
	fmt.Printf("test: get() -> [status:%v] [path:%v] [entries:%v]\n", status, path, len(entries))

	//Output:
	//test: get() -> [status:OK] [path:storage/address?state=IA] [entries:2]

}
