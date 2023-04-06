package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/raphael-foliveira/simple-chat/backend/models"
)

var Db *sql.DB

func GetDb() {
	var err error
	Db, err = sql.Open("postgres", "postgresql://postgres:postgres@database:5432/local-chat?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func SaveMessage(message models.Message) error {
	_, err := Db.Exec("INSERT INTO messages (sender, content, chat_name, sent_at) VALUES ($1, $2, $3, $4)", message.Sender, message.Content, message.ChatName, message.SentAt)
	return err
}

func GetChatRoomMessages(chatName string) ([]models.Message, error) {
	rows, err := Db.Query("SELECT id, sender, content, chat_name, sent_at FROM messages WHERE chat_name = $1", chatName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	messages := []models.Message{}
	for rows.Next() {
		var message models.Message
		err = rows.Scan(&message.Id, &message.Sender, &message.Content, &message.ChatName, &message.SentAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, err
}
