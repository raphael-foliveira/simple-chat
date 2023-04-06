package models

import "time"

type Message struct {
	Id       uint      `json:"id"`
	Sender   string    `json:"sender"`
	Content  string    `json:"content"`
	ChatName string    `json:"chatName"`
	SentAt   time.Time `json:"sentAt"`
}
