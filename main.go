package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, you are visiting %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	port := 8080
	log.Printf("Starting server %d \n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}
