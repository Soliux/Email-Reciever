package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebsocket(conn *websocket.Conn) {
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Failed to read WebSocket message: %v\n", err)
			return
		}

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Printf("Failed to write WebSocket message: %v\n", err)
			return
		}
	}
}

func WebsocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade connection to WebSocket: %v\n", err)
		return
	}

	handleWebsocket(conn)
}
