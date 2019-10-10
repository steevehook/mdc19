package controllers

import (
	"encoding/json"
	"net/http"
)

type commentsService interface {
	GetComments() ([]string, error)
}

func GetComments(service commentsService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			comments, err := service.GetComments()
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			err = json.NewEncoder(w).Encode(comments)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
