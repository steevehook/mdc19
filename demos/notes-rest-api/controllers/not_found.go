package controllers

import (
	"net/http"

	"github.com/steevehook/mdc19/dummy-notes-rest-api/models"
	"github.com/steevehook/mdc19/dummy-notes-rest-api/transport"
)

func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		transport.SendError(w, models.NotFoundError{})
	})
}
