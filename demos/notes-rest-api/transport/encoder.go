package transport

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"

	"github.com/steevehook/mdc19/notes-rest-api/logging"
	"github.com/steevehook/mdc19/notes-rest-api/models"
)

// HTTP constants
const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
)

func toHTTPError(err error) models.HTTPError {
	httpErr := models.HTTPError{Message: err.Error()}
	switch err.(type) {
	case models.NotFoundError:
		httpErr.Type = models.NotFoundErrorType
		httpErr.StatusCode = http.StatusNotFound
	case models.DataValidationError:
		httpErr.Type = models.DataValidationErrorType
		httpErr.StatusCode = http.StatusBadRequest
	case models.HTTPError:
		httpErr.Type = models.HTTPErrorType
		httpErr.StatusCode = http.StatusInternalServerError
	default:
		httpErr.Type = models.HTTPErrorType
		httpErr.StatusCode = http.StatusInternalServerError
	}
	return httpErr
}

func jsonEncoder(w http.ResponseWriter, statusCode int) *json.Encoder {
	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w)
}

func SendJSON(w http.ResponseWriter, code int, data interface{}) {
	err := jsonEncoder(w, code).Encode(data)
	if err != nil {
		logging.Logger.With(zap.Error(err))
	}
}

func SendError(w http.ResponseWriter, err error) {
	httpErr := toHTTPError(err)
	SendJSON(w, httpErr.StatusCode, httpErr)
}

func SendPanicError(w http.ResponseWriter, e error) {
	err := toHTTPError(e)
	logging.Logger.Error("caught unexpected panic", zap.Error(err), zap.Stack("stack trace"))
	SendError(w, err)
}
