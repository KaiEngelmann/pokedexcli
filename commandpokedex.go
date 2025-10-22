package main

import "fmt"

func commandPokedex(cfg *Config, args []string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Printf("You haven't caught a Pokemon yet\n")
		return nil
	}
	fmt.Printf("Your Pokedex\n")
	for _, entries := range cfg.Pokedex {
		fmt.Printf(" - %s\n", entries.pokemon.Name)
	}
	return nil
}
