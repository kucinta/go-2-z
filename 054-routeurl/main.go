package main

import (
	"log"
	"net/http"
)

//import "fmt"

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/about", about)
	http.HandleFunc("/error", http.NotFound)
	log.Println("Listening on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	println("This is a " + r.Method + " request")
	//fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	url := r.URL.String()
	println(url)
	if url == "/" {
		w.Write([]byte("Hello World"))
	} else {
		w.Write([]byte("404 " + url + " page not found"))
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me!"))
}
