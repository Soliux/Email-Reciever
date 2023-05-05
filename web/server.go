package web

import (
	"email-reciever/database"
	"email-reciever/web/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	DatabaseConnect()
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

func DatabaseConnect() {
	if err := database.ConnectDB(); err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	defer database.DB.Close()

	if err := database.DB.DB().Ping(); err != nil {
		fmt.Println("Failed to ping database:", err)
		return
	}

	fmt.Println("Connected to database!")
}
