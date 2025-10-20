package main

import "fmt"

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location area name, e.g., 'canalave-city-area'")
	}

	area := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)

	exploredArea, err := cfg.PokeClient.ExploreArea(url)
	if err != nil {
		return fmt.Errorf("failed to explore area: %w", err)
	}

	fmt.Printf("Exploring area: %s (ID: %d)\n", exploredArea.Name, exploredArea.ID)

	for _, encounter := range exploredArea.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
