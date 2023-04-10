package handlers

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/raphael-foliveira/simple-chat/backend/database"
	"github.com/raphael-foliveira/simple-chat/backend/models"
)

type ChatSubscriber struct {
	Conn *websocket.Conn
}

var rooms = struct {
	sync.RWMutex
	m map[string][]ChatSubscriber
}{m: make(map[string][]ChatSubscriber)}

func WsServer(c *websocket.Conn) {
	defer c.Close()
	fmt.Println("running websocket")

	chatName := c.Params("chatName")

	newSubscriber := ChatSubscriber{Conn: c}
	rooms.Lock()
	_, exists := rooms.m[chatName]
	if !exists {
		fmt.Println("creating new chat:", chatName)
		rooms.m[chatName] = []ChatSubscriber{}
	}
	fmt.Println("adding new subscriber to", chatName)
	rooms.m[chatName] = append(rooms.m[chatName], newSubscriber)

	rooms.Unlock()

	for {
		var message models.Message
		err := c.ReadJSON(&message)
		if err != nil {
			fmt.Println("Error in receiving:", err)
			break
		}
		database.SaveMessage(&message)
		for _, sub := range rooms.m[chatName] {
			sub.Conn.WriteJSON(message)
		}
		fmt.Println("Received from", message.Sender, "in", chatName+":", message.Content)
	}
	delete(rooms.m, chatName)
}
