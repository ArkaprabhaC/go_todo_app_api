package errors

import (
	"net/http"
	"time"
)

type AppError struct {
	Status string
	Timestamp time.Time
	Code int
}

func NewRequestTimeoutError() *AppError {
	return &AppError{
		Status: http.StatusText(http.StatusRequestTimeout),
		Code: http.StatusRequestTimeout,
		Timestamp: time.Now().UTC(),
	}
}

var REQUEST_BODY_PARSE_ERROR = AppError{
	Status: http.StatusText(http.StatusBadRequest),
	Code: http.StatusBadRequest,
	Timestamp: time.Now().UTC(),
}

var FAILURE_TO_ADD_NOTE_ERROR = AppError {
	Status: http.StatusText(http.StatusInternalServerError),
	Code: http.StatusInternalServerError,
	Timestamp: time.Now().UTC(),
}