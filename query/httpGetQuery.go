package query

import (
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	errorDuringGetRequest = errors.New("get request failed")
)

// HTTPGetQuery is a basic client to make the query
type HTTPGetQuery struct{}

// Get is a function that performs basic GET reqests
func (HTTPGetQuery) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errorDuringGetRequest
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
