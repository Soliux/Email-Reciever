package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func newMessage(c *gin.Context) {
	domain := c.Query("domain")
	email := c.Query("email")

	if domain == "" || email == "" {
		c.JSON(400, gin.H{
			"message": "Missing domain or email",
		})
		return
	}

	fmt.Println("New Message Recieved!")
	fmt.Println("Domain: " + domain)
	fmt.Println("Email: " + email)

	c.JSON(200, gin.H{
		"message": "Our New Message Route",
	})
}
