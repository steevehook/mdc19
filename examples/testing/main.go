package main

import (
	"log"
	"net/http"

	"github.com/steevehook/mdc19/examples/testing/controllers"
	"github.com/steevehook/mdc19/examples/testing/services"
)

func main() {
	commentsService := services.NewComments()
	http.HandleFunc(
		"/comments",
		controllers.GetComments(commentsService),
	)
	log.Fatal(http.ListenAndServe(
		":8080",
		nil,
	))
}
