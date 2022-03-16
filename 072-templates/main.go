package main

import (
  "html/template"
  "net/http"
  "fmt"
)

type ViewData struct {
  Name string
  Ic string
  Items []Item
}

type Item struct {
	Name  string
	Price int
}

var myTmplvar = `
{{define "footer"}}
  <footer>
    <p>
      My Custom Footer!
    </p>
  </footer>
{{end}}
<h1>{{if .Name}} {{.Name}}-{{.Ic}} {{else}} there {{end}}</h1>
<p>
{{range .Items}}
    {{.Name}} = ${{.Price}}<br/>
  </div>
{{end}}
</p>
{{template "footer"}}
`

func handler(w http.ResponseWriter, r *http.Request) {
	testTemplate1,err := template.New("").Parse(myTmplvar)
		if err != nil {
			panic(err)
		}
	w.Header().Set("Content-Type", "text/html")
	vd := ViewData{Name: "John", Ic: "123",Items: []Item{Item{"Item No 1",100},Item{"Item No 2",123}}}
	err2 := testTemplate1.Execute(w, vd)
	if err2 != nil {
	http.Error(w, err2.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	myPort := ":8072"
	fmt.Println( "Running on 192.168.1.234" + myPort)
	http.ListenAndServe(myPort, nil)
  }
