package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func main() {
	input := []byte("# 123\n123")
	unsafe := blackfriday.Run(input)
	myHTMLbyte := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	myHTMLstring := string(myHTMLbyte)
	fmt.Println(myHTMLstring)
}
