package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// New function returns the router for the whole project
// with the initialized handlers
func New() *gin.Engine {
	router := gin.Default()

	router.GET("/all", getAll)
	router.GET("/topic", getTopic)
	router.POST("/topic", requestTopic)

	return router
}

func getAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func getTopic(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "topic"})
}

func requestTopic(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "requested"})
}
