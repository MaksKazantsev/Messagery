package models

import "time"

type Message struct {
	Value    string    `json:"value"`
	SendAt   time.Time `json:"sendAt"`
	SenderID string    `json:"senderID"`
}
