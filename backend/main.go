package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatSubscriber struct {
	Conn *websocket.Conn
}

var subscribers = struct {
	sync.RWMutex
	m []ChatSubscriber
}{}

func WsServer(ws *websocket.Conn) {
	defer ws.Close()

	newSubscriber := ChatSubscriber{Conn: ws}
	subscribers.Lock()
	subscribers.m = append(subscribers.m, newSubscriber)
	subscribers.Unlock()

	for {
		var message Message
		err := ws.ReadJSON(&message)
		if err != nil {
			fmt.Println("Error in receiving:", err)
			break
		}
		for _, subscriber := range subscribers.m {
			subscriber.Conn.WriteJSON(message)
		}

		fmt.Println("Received from", message.Sender, ":", message.Content)
	}
}

func main() {
	app := fiber.New()

	app.Get("/chat", websocket.New(WsServer))

	app.Listen(":8000")
}
