package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	stdEncoding := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(stdEncoding)

	stdDecoding, _ := b64.StdEncoding.DecodeString(stdEncoding)
	fmt.Println(string(stdDecoding), "\n")

	urlEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEnc)

	urlDec, _ := b64.URLEncoding.DecodeString(urlEnc)
	fmt.Println(string(urlDec))
}