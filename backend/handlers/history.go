package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/simple-chat/backend/database"
)

func GetHistory(c *fiber.Ctx) error {
	messages, err := database.GetChatRoomMessages(c.Params("chatName"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.Status(200).JSON(messages)
}
