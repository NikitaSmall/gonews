package models

// NewsResponse is an object that's returned from News API global request
type NewsResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}
