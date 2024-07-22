package dto

import (
	"time"
)

type CreateNoteRequest struct {
	CorrelationId    string
	RequestTimestamp time.Time
	Title            string
	Description      string
}
