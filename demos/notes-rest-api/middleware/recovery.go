package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/steevehook/mdc19/dummy-notes-rest-api/models"
)

func Recovery(sendError models.ErrorHandlerFunc) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					var panicErr error
					switch x := err.(type) {
					case string:
						panicErr = errors.New(x)
					case error:
						panicErr = x
					default:
						panicErr = fmt.Errorf("unknown panic: %v", x)
					}
					sendError(w, panicErr)
				}
			}()

			h.ServeHTTP(w, r)
		})
	}
}
