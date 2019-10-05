package flow

import (
	"log"

	"github.com/nikitasmall/gonews/models"
)

type ChanHandlerFlow struct {
	RequestNewsChan chan string
	RequestChan     chan string
	DataChan        chan []models.Article

	GetNewsDataFlow
}

func (flow ChanHandlerFlow) Handle() {
	dataMap := make(map[string][]models.Article)

	for {
		select {
		case topic := <-flow.RequestChan:
			articles, err := flow.GetNewsDataFlow.GetData(topic)
			if err != nil {
				log.Println(err)
				continue
			}

			dataMap[topic] = articles
		case topic := <-flow.RequestNewsChan:
			data, ok := dataMap[topic]
			if !ok {
				flow.DataChan <- make([]models.Article, 0)
				continue
			}

			flow.DataChan <- data
		}
	}
}
