package flow

import (
	"log"

	"github.com/nikitasmall/gonews/config"
	"github.com/nikitasmall/gonews/models"
	"github.com/nikitasmall/gonews/query"
)

type ChanHandlerFlow struct {
	RequestNewsChan chan string
	RequestChan     chan string
	DataChan        chan []models.Article
}

func (flow ChanHandlerFlow) Handle() {
	newsDataFlow := GetNewsDataFlow{
		HTTPGetQuery: query.HTTPGetQuery{},
		NewsAPIURLQuery: query.NewNewsAPIURLQuery(
			config.NewsAPIDomain, config.NewsAPIKey),
	}

	dataMap := make(map[string][]models.Article)

	for {
		select {
		case topic := <-flow.RequestChan:
			articles, err := newsDataFlow.GetData(topic)
			if err != nil {
				log.Println(err)
				continue
			}

			dataMap[topic] = articles
		case topic := <-flow.RequestNewsChan:
			flow.DataChan <- dataMap[topic]
		}
	}
}
