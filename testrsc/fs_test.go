package testrsc

import (
	"fmt"
	"github.com/advanced-go/common/iox"
)

func ExampleReadFile() {
	name := "file:///f:/files/address1/entry.json"
	bytes, status := iox.ReadFile(name)
	fmt.Printf("test: ReadFile() -> [buff:%v] [status:%v]\n", len(bytes), status)

	//Output:
	//test: ReadFile() -> [buff:750] [status:OK]

}
