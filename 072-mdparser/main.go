package main

import (
	"fmt"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func main() {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// md is the markdown string
	//myMDstring := "# Intro\ntest\n```\nthis code block will be dropped from output\n```\ntext"
	myMDstring := "# Intro\ntest\n```\ncode\n```\ntest.\n123"

	// need to use the []byte() function to convert to byte for processing
	myHTMLbyte := markdown.ToHTML([]byte(myMDstring), nil, renderer)
	// output also in bytehtml
	myHTMLstring := string(myHTMLbyte)
	fmt.Println(myHTMLstring)
}
