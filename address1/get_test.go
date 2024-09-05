package address1

import (
	"fmt"
	"github.com/advanced-go/customer/testrsc"
	"github.com/advanced-go/stdlib/core"
	"net/url"
)

func ExampleGet() {
	ex := core.NewExchangeOverride("", testrsc.Addr1EntryURL, "")
	ctx := core.NewExchangeOverrideContext(nil, ex)
	values := make(url.Values)

	values.Add(core.RegionKey, "us-west")
	entries, _, status := get[core.Output](ctx, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Add(core.SubZoneKey, "dc1")
	entries, _, status = get[core.Output](ctx, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}
