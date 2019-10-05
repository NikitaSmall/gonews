package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/gonews/config"
	"github.com/nikitasmall/gonews/flow"
	"github.com/nikitasmall/gonews/models"
	"github.com/nikitasmall/gonews/query"
)

// New function returns the router for the whole project
// with the initialized handlers
func New() *gin.Engine {
	router := gin.Default()

	newsChannel := make(chan []models.Article)
	requestTopicChannel := make(chan string)
	requestDataChannel := make(chan string)

	newsDataFlow := flow.GetNewsDataFlow{
		HTTPGetQuery: query.HTTPGetQuery{},
		NewsAPIURLQuery: query.NewNewsAPIURLQuery(
			config.NewsAPIDomain, config.NewsAPIKey),
	}

	backgroundProcessor := flow.ChanHandlerFlow{
		RequestChan:     requestTopicChannel,
		RequestNewsChan: requestDataChannel,
		DataChan:        newsChannel,
		GetNewsDataFlow: newsDataFlow,
	}

	go backgroundProcessor.Handle()

	router.GET("/all", getAll)
	router.GET("/topic/:topic", getTopic(requestDataChannel, newsChannel))
	router.POST("/topic/:topic", requestTopic(requestTopicChannel))

	return router
}

func getAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func getTopic(requestDataChannel chan string, dataChan chan []models.Article) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		topic := ctx.Param("topic")

		requestDataChannel <- topic

		ctx.JSON(http.StatusOK, <-dataChan)
	}
}

func requestTopic(requestTopicChannel chan string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		topic := ctx.Param("topic")

		requestTopicChannel <- topic

		ctx.JSON(http.StatusOK, gin.H{"status": "requested"})
	}
}
