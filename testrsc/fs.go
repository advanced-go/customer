package testrsc

import (
	"embed"
	"github.com/advanced-go/stdlib/io"
)

//go:embed files
var f embed.FS

func init() {
	io.Mount(f)
}

const (
	Addr1BasePath = "file:///f:/files/address1"

	Addr1EntryURL = Addr1BasePath + "/entry.json"

	Addr1GetReqURL  = Addr1BasePath + "/get-req.txt"
	Addr1GetRespURL = Addr1BasePath + "/get-resp.txt"

	Addr1GetReqErrHeaderURL   = Addr1BasePath + "/get-req-error-header.txt"
	Addr1GetRespErrContextURL = Addr1BasePath + "/get-resp-error-content.txt"
)
