package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/simple-chat/backend/database"
)

func GetHistory(c *fiber.Ctx) error {
	messages := database.GetChatRoomMessages(c.Params("chatName"))
	fmt.Println("messages:", messages)
	return c.Status(200).JSON(messages)
}
