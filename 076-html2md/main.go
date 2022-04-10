package main

import (
	"fmt"
	"log"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	converter := md.NewConverter("", true, nil)
	html := `<p class="fr-tag">
	<img class="right fr-tag fr-fir" alt="/im/dbedu/1586487.jpeg" src="/im/dbedu/1586487.jpeg">
	</p><ul class="fr-tag"><li class="fr-tag">Choose <a href="139820.page">RPN</a>`
	//htmlb := []byte(html)
	//htmlout := bytes.Replace(htmlb, []byte(".page"), []byte("/"), -1)
	//html = string(htmlout)
	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdown)
	temp := strings.Replace(markdown, ".page", "/", -1)
	temp = strings.Replace(temp, "](", "(/page/", -1)
	markdownout := strings.Replace(temp, "/page//", "/", -1)
	fmt.Println(markdownout)

}
