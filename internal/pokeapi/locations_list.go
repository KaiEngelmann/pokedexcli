package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kaiengelmann/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	cache      *pokecache.Cache
}

func NewClient(httpClient *http.Client, baseURL string, cache *pokecache.Cache) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		cache:      cache,
	}
}

func (c *Client) ListLocationAreas(url string) (LocationAreaList, error) {
	var locationAreas LocationAreaList
	if entry, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(entry, &locationAreas)
		return locationAreas, err
	}
	if url == "" {
		url = c.baseURL + "/location-area/"
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return locationAreas, fmt.Errorf("failed to get URL %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		body, _ := io.ReadAll(res.Body)
		return locationAreas, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, string(body))
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas, fmt.Errorf("failed to read body: %w", err)
	}

	if err := json.Unmarshal(body, &locationAreas); err != nil {
		return locationAreas, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return locationAreas, nil
}
