package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	fmt.Println("Web Server Is Online!")
	fmt.Println("Web Socket Server Is Online!")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	websocket := api.Group("/ws")

	api.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	websocket.GET("/", websocketHandler)
	r.Run()
}

func ShutdownServer() {
	fmt.Println("Web Server Is Offline!")
	fmt.Println("Web Socket Server Is Offline!")
}
