package model

import (
	"time"
)

type Message struct {
	IsSender bool      `json:"isSender"`
	SentAt   time.Time `json:"sent_at"`
	Message  string    `json:"message"`
}
