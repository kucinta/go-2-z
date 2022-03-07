package main

import (
	"065-deploy/handlers"
	"log"
	"net/http"
	"fmt"
	
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Get("/api/j", handlers.GetJobs)
	fmt.Println("192.168.1.234:8065")
	err := http.ListenAndServe("0.0.0.0:8065", router)
	if err != nil {
		log.Fatal(err)
	}
}
