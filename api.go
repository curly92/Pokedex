package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationResponse struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func request(url string) (LocationResponse, error) {
	location := LocationResponse{}
	res, err := http.Get(url)
	if err != nil {
		return location, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return location, err
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		return location, fmt.Errorf("response failed with status code: %v", res.StatusCode)
	}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return location, err
	}
	return location, nil

}
