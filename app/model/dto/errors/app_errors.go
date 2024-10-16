package errors

import (
	"net/http"
	"time"
)

type AppError struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Code      int       `json:"code"`
	Message   string    `json:"message"`
}

func NewRequestTimeoutError() *AppError {
	return &AppError{
		Status:    http.StatusText(http.StatusRequestTimeout),
		Code:      http.StatusRequestTimeout,
		Timestamp: time.Now().UTC(),
		Message:   "Request timed out",
	}
}

var REQUEST_BODY_PARSE_ERROR = AppError{
	Status:    http.StatusText(http.StatusBadRequest),
	Code:      http.StatusBadRequest,
	Timestamp: time.Now().UTC(),
	Message:   "Unable to parse request body",
}

var FAILURE_TO_ADD_NOTE_ERROR = AppError{
	Status:    http.StatusText(http.StatusInternalServerError),
	Code:      http.StatusInternalServerError,
	Timestamp: time.Now().UTC(),
	Message:   "Unable to add a note in the system",
}
