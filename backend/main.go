package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
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
		err := websocket.JSON.Receive(ws, &message)
		if err != nil {
			fmt.Println("Error in receiving:", err)
			break
		}
		for _, subscriber := range subscribers.m {
			websocket.JSON.Send(subscriber.Conn, message)
		}

		fmt.Println("Received from", message.Sender, ":", message.Content)
	}
}

func main() {
	router := gin.Default()

	router.GET("/chat", func(c *gin.Context) {
		handler := websocket.Handler(WsServer)
		handler.ServeHTTP(c.Writer, c.Request)
	})

	router.Run(":8000")
}
