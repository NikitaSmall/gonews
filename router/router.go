package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/gonews/config"
	"github.com/nikitasmall/gonews/flow"
	"github.com/nikitasmall/gonews/query"
)

// New function returns the router for the whole project
// with the initialized handlers
func New() *gin.Engine {
	router := gin.Default()

	router.GET("/all", getAll)
	router.GET("/topic", getTopic)
	router.POST("/topic/:topic", requestTopic)

	return router
}

func getAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func getTopic(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "topic"})
}

func requestTopic(ctx *gin.Context) {
	topic := ctx.Param("topic")

	newsDataFlow := flow.GetNewsDataFlow{
		HTTPGetQuery: query.HTTPGetQuery{},
		NewsAPIURLQuery: query.NewNewsAPIURLQuery(
			config.NewsAPIDomain, config.NewsAPIKey),
	}

	data, err := newsDataFlow.GetData(topic)
	if err != nil {
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}
