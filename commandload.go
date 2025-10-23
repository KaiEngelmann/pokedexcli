package main

import (
	"fmt"
	"os"
	"strings"
)

func commandLoad(cfg *Config, args []string) error {
	saved_name := ""
	if len(args) > 0 {
		saved_name = args[0]
	}
	filename := "pokedex_entry_" + saved_name + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to load pokedex")
		return err
	}
	fileContent := string(data)
	rawNames := strings.Split(fileContent, "\n")
	var caughtPokemon []string
	for _, name := range rawNames {
		trimmedName := strings.TrimSpace(name)
		if trimmedName != "" {
			caughtPokemon = append(caughtPokemon, trimmedName)
		}
	}
	for _, name := range caughtPokemon {
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
		pokemon, err := cfg.PokeClient.CheckPokemon(url)
		if err != nil {
			fmt.Printf("Failed to load %s: %v", name, err)
			continue
		}
		cfg.Pokedex = append(cfg.Pokedex, Pokedex{pokemon: pokemon})
	}
	fmt.Println("Sucessfully loaded all Pokémon from file.")
	return nil
}
