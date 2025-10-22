package main

import (
	"fmt"
	"strings"

	"github.com/kaiengelmann/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a Pokémon name")
	}

	name := args[0]
	for _, entry := range cfg.Pokedex {
		if entry.pokemon.Name == name {
			name := entry.pokemon.Name
			height := entry.pokemon.Height
			weight := entry.pokemon.Weight
			hp := getAdvancedStats(entry.pokemon, "hp")
			attack := getAdvancedStats(entry.pokemon, "attack")
			defense := getAdvancedStats(entry.pokemon, "defense")
			special_attack := getAdvancedStats(entry.pokemon, "special-attack")
			special_defense := getAdvancedStats(entry.pokemon, "special-defense")
			speed := getAdvancedStats(entry.pokemon, "speed")
			types := getTypes(entry.pokemon)

			fmt.Printf("Found Pokémon: %s\n", name)
			fmt.Printf("Height: %d\n", height)
			fmt.Printf("Weight: %d\n", weight)
			fmt.Printf("Stats:\n")
			fmt.Printf("  -hp: %d\n", hp)
			fmt.Printf("  -attack: %d\n", attack)
			fmt.Printf("  -defense: %d\n", defense)
			fmt.Printf("  -special-attack: %d\n", special_attack)
			fmt.Printf("  -special-defense: %d\n", special_defense)
			fmt.Printf("  -speed: %d\n", speed)
			fmt.Printf("Type(s): %s\n", strings.Join(types, ", "))
			return nil
		}
	}

	fmt.Println("You haven't caught that Pokémon yet.")
	return nil
}

func getAdvancedStats(p pokeapi.Pokemon, statName string) int {
	for _, s := range p.Stats {
		if s.Stat.Name == statName {
			return s.BaseStat
		}
	}
	return 0
}

func getTypes(p pokeapi.Pokemon) []string {
	var types []string
	for _, t := range p.Types {
		types = append(types, t.Type.Name)

	}
	return types
}
