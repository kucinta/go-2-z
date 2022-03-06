package main

import (
	"log"
	"net/http"
)

//import "fmt"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	//http.HandleFunc("/", hello)
	http.HandleFunc("/about", about)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/every-thing-else", http.NotFound)
	log.Println("Listening on http://0.0.0.0:8054")
	log.Fatal(http.ListenAndServe(":8054", nil))
}

func about(w http.ResponseWriter, r *http.Request) {
	println("This is a " + r.Method + " request")
	url := r.URL.String()
	println(url)
	w.Write([]byte("About me!"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there!"))
}
