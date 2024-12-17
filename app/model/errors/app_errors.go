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

func (e *AppError) Error() string {
	return e.Message
}

var FailureBadRequest = AppError{
	Status:    http.StatusText(http.StatusBadRequest),
	Code:      http.StatusBadRequest,
	Timestamp: time.Now().UTC(),
	Message:   "Unable to parse request",
}

var FailureAddNoteError = AppError{
	Status:    http.StatusText(http.StatusInternalServerError),
	Code:      http.StatusInternalServerError,
	Timestamp: time.Now().UTC(),
	Message:   "Unable to add a note in the system",
}

var FailureGetNotes = AppError{
	Status:    http.StatusText(http.StatusInternalServerError),
	Code:      http.StatusInternalServerError,
	Timestamp: time.Now().UTC(),
	Message:   "Unable to get notes from the system",
}

var FailureNoteAlreadyExists = AppError{
	Status:    http.StatusText(http.StatusInternalServerError),
	Code:      http.StatusInternalServerError,
	Timestamp: time.Now().UTC(),
	Message:   "Note with same title already exists. Please use a different title and try again",
}

var FailureNoteNotFound = AppError{
	Status:    http.StatusText(http.StatusNotFound),
	Code:      http.StatusNotFound,
	Timestamp: time.Now().UTC(),
	Message:   "Note with given title is not found",
}
