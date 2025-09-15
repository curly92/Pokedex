package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/curly92/Pokedex/internal/pokecache"
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

type Client struct {
	http  *http.Client
	cache *pokecache.Cache
}

func NewClient(c *pokecache.Cache) *Client {
	return &Client{
		http:  &http.Client{Timeout: 10 * time.Second},
		cache: c,
	}
}

func (client *Client) fetch(url string) ([]byte, error) {
	if b, ok := client.cache.Get(url); ok {
		fmt.Println("CACHE WAS USED")
		return b, nil
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("bad status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	client.cache.Add(url, body)
	fmt.Println("CACHE WAS NOT USED")
	return body, nil
}

func (c *Client) GetLocations(url string) (LocationResponse, error) {
	var locations LocationResponse

	data, err := c.fetch(url)
	if err != nil {
		return locations, err
	}
	if err := json.Unmarshal(data, &locations); err != nil {
		return locations, err
	}
	return locations, nil

}
