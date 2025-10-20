package main

import (
	"fmt"
)

func commandMapb(cfg *Config, args []string) error {
	url := ""
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url = *cfg.Previous
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
