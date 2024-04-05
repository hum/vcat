package vcat

import (
	"io"
	"net/http"
	"time"
)

var httpclient *http.Client = &http.Client{Timeout: 60 * time.Second}

func do(httpclient *http.Client, url string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpclient.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
