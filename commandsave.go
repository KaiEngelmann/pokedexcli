package main

import (
	"fmt"
	"os"
	"strings"
)

func commandSave(cfg *Config, args []string) error {
	var pokemon_names []string
	for _, entries := range cfg.Pokedex {
		pokemon_names = append(pokemon_names, entries.pokemon.Name)
	}
	data := []byte(strings.Join(pokemon_names, "\n"))
	filename := "pokedex_entry.txt"
	permissions := os.FileMode(0644)
	err := os.WriteFile(filename, data, permissions)
	if err != nil {
		fmt.Printf("Oh no! Failed to save to file")
		return err
	}

	fmt.Printf("Successfully saved game\n")
	return nil
}
