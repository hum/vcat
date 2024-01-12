package main

import (
	"io"
	"net/http"
)

// GetBodyAsByteSlice makes an HTTP request to the provided URL
// and returns the response's body as a byte slice.
func GetBodyAsByteSlice(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	return body, err
}
