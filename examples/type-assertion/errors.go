package main

const (
	HTTPErrorType      = "http_error"
	DataValidationType = "data_validation"
)

type DataValidationError struct {
	Type        string                 `json:"type"`
	Message     string                 `json:"message"`
	Constraints map[string]interface{} `json:"constraints"`
}

func (e DataValidationError) Error() string {
	return e.Message
}

type HTTPError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}
