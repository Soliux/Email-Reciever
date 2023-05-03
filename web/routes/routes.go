package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	registerMainWebRoutes(router)
	registerAPIRoutes(router)
	registerWebsocketRoutes(router)
}

func registerAPIRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")

	apiGroup.GET("/message", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is an example API route",
		})
	})
}

func registerWebsocketRoutes(router *gin.Engine) {
	router.GET("/ws", WebsocketHandler)
}

func registerMainWebRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
