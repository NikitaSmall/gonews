package query

import "fmt"

// NewsAPIURLQuery holds the data in order to create correct URL
type NewsAPIURLQuery struct {
	Domain string
	Query  string
	APIKey string
}

// NewNewsAPIURLQuery creates an object that is used as a URL builder
func NewNewsAPIURLQuery(domain, apiKey string) NewsAPIURLQuery {
	return NewsAPIURLQuery{
		Domain: domain,
		APIKey: apiKey,
	}
}

// GetURL returns a ready to use URL with correct params
func (query NewsAPIURLQuery) GetURL(queryParam string) string {
	return fmt.Sprintf("%s/v2/everything?q=%s&apiKey=%s",
		query.Domain, queryParam, query.APIKey)
}
