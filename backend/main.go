package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"github.com/raphael-foliveira/simple-chat/backend/database"
	"github.com/raphael-foliveira/simple-chat/backend/handlers"
)

func main() {
	database.GetDb()
	defer database.Db.Close()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/chat/:chatName", websocket.New(handlers.WsServer))
	app.Get("/messages/:chatName", handlers.GetHistory)

	app.Listen(":8000")
}
