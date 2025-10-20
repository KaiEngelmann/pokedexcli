package main

import (
	"fmt"
)

func commandMap(cfg *Config, args []string) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
	}
	resp, err := cfg.PokeClient.ListLocationAreas(url)
	if err != nil {
		return err
	}
	for _, r := range resp.Results {
		fmt.Println(r.Name)
	}
	if resp.Next != nil {
		cfg.Next = resp.Next
	} else {
		cfg.Next = nil
	}
	if resp.Previous != nil {
		cfg.Previous = resp.Previous
	} else {
		cfg.Previous = nil
	}
	return nil
}
