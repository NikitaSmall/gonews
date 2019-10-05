package models

// Article is a main object that is returned from API
type Article struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
	Desc   string `json:"description"`

	Topic string `json:"-"`
}
