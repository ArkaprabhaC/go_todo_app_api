package errors

import (
	"net/http"
	"time"
)

type RequestTimeoutError struct {
	Status string
	Timestamp time.Time
	Code int
}

func NewRequestTimeoutError() *RequestTimeoutError {
	return &RequestTimeoutError{
		Status: http.StatusText(http.StatusRequestTimeout),
		Code: http.StatusRequestTimeout,
		Timestamp: time.Now().UTC(),
	}
}