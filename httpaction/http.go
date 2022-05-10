package httpaction

import (
	"io/ioutil"
	"net/http"
)

type HTTPAction struct{}
type ConcurrentResult struct {
	Result []byte
	Error  error
}

func (httpaction *HTTPAction) Get(url string) ([]byte, error) {
	var body []byte
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	return body, err
}

func (httpaction *HTTPAction) GetCuncurrency(url string, retryCount int, ch chan<- ConcurrentResult) {
	var resp []byte
	var err error
	for retryCount > 0 {
		resp, err = httpaction.Get(url)
		if err != nil {
			retryCount -= 1
		} else {
			break
		}
	}
	ch <- ConcurrentResult{
		Result: resp,
		Error:  err,
	}
}
