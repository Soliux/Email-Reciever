package web

import (
	"email-reciever/web/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	fmt.Println("Web Server Is Online!")
	fmt.Println("Web Socket Server Is Online!")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	r.Run()
}

func ShutdownServer() {
	fmt.Println("Web Server Is Offline!")
	fmt.Println("Web Socket Server Is Offline!")
}
