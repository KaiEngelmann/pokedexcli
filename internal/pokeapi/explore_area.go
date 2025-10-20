package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ExploreArea(url string) (ExploreArea, error) {
	var exploredArea ExploreArea
	if entry, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(entry, &exploredArea)
		return exploredArea, err
	}
	if url == "" {
		url = c.baseURL + "/location-area/"
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return exploredArea, fmt.Errorf("failed to get URL %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		body, _ := io.ReadAll(res.Body)
		return exploredArea, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, string(body))
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return exploredArea, fmt.Errorf("failed to read body: %w", err)
	}

	if err := json.Unmarshal(body, &exploredArea); err != nil {
		return exploredArea, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return exploredArea, nil
}
