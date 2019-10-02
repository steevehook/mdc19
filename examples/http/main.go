package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home())
	mux.HandleFunc("/about", about())
	log.Fatal(http.ListenAndServe(":8080", mux))
}
