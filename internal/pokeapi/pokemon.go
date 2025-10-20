package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) CheckPokemon(url string) (Pokemon, error) {
	var pokemon Pokemon
	if entry, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(entry, &pokemon)
		return pokemon, err
	}
	if url == "" {
		url = c.baseURL + "/pokemon/"
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return pokemon, fmt.Errorf("failed to get URL %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		body, _ := io.ReadAll(res.Body)
		return pokemon, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, string(body))
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemon, fmt.Errorf("failed to read body: %w", err)
	}

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return pokemon, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return pokemon, nil
}
