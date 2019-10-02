package main

import (
	"fmt"
	"net/http"
)

func home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handle errors, son of the bitch
		_, _ = fmt.Fprint(w, "Welcome Home!")
	}
}
