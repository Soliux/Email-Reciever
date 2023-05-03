package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	fmt.Println("Web Server Is Online!")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run()
}

func ShutdownServer() {
	fmt.Println("Web Server Is Offline!")
}
