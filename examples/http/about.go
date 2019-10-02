package main

import (
	"fmt"
	"net/http"
)

func about() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handle errors, son of the bitch
		_, _ = fmt.Fprint(w, "About")
	}
}
