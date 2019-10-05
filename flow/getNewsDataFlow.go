package flow

import (
	"encoding/json"

	"github.com/nikitasmall/gonews/models"
	"github.com/nikitasmall/gonews/query"
)

type GetNewsDataFlow struct {
	query.HTTPGetQuery
	query.NewsAPIURLQuery
}

func (flow *GetNewsDataFlow) GetData() ([]models.Article, error) {
	url := flow.NewsAPIURLQuery.GetURL()

	body, err := flow.HTTPGetQuery.Get(url)
	if err != nil {
		return nil, err
	}

	var newsAPIResponse models.NewsResponse
	if err := json.Unmarshal(body, &newsAPIResponse); err != nil {
		return nil, err
	}

	return newsAPIResponse.Articles, nil
}
