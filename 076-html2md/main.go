package main
import (
	"fmt"
	"log"
	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {
	converter := md.NewConverter("", true, nil)
	html := `<html><h1>hello</h1><h3>world</h3><p>This is <strong>Important</strong></p><p>The End</p></html>`
	markdown, err := converter.ConvertString(html)
	if err != nil {
	  log.Fatal(err)
	}
	fmt.Println(markdown)	
}

