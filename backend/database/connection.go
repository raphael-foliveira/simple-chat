package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/raphael-foliveira/simple-chat/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetDb() {
	var err error
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	Db, err = gorm.Open(postgres.Open(dbUrl))
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&models.Message{})
	fmt.Println("database connected")
}

func SaveMessage(message *models.Message) {
	Db.Create(message)
}

func GetChatRoomMessages(chatName string) []models.Message {
	var messages []models.Message
	Db.Where(models.Message{ChatName: chatName}).Find(&messages)
	return messages
}
