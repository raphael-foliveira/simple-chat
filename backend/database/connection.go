package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/raphael-foliveira/simple-chat/backend/models"
)

var Db *sql.DB

func GetDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	Db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected")
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
