package model

import (
	"time"
)

type Room struct {
	RoomID     string    `json:"room_id"`
	ReceiverID string    `json:"receiver_user_id"`
	SenderID   string    `json:"sender_user_id"`
	Messages   []Message `json:"messages"`
	CreatedAt  time.Time `json:"created_at"`
}
