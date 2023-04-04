package main

import "time"

type Message struct {
	Id      string    `json:"id"`
	Sender  string    `json:"sender"`
	Content string    `json:"content"`
	SentAt  time.Time `json:"sentAt"`
}
