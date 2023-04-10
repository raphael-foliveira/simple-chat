package models

import "time"

type Message struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
	ChatName  string    `json:"chatName"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
