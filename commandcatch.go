package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kaiengelmann/pokedexcli/internal/pokeapi"
)

func getStats(cfg *Config, args []string) (pokeapi.Pokemon, int, error) {
	if len(args) == 0 {
		return pokeapi.Pokemon{}, 1, fmt.Errorf("please provide a pokemon to catch")
	}

	pokeName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokeName)

	attemptPokemon, err := cfg.PokeClient.CheckPokemon(url)
	if err != nil {
		return pokeapi.Pokemon{}, 1, fmt.Errorf("failed to find Pokemon")
	}

	hp := 0
	for _, s := range attemptPokemon.Stats {
		if s.Stat.Name == "hp" {
			hp = s.BaseStat
			break
		}
	}

	return attemptPokemon, hp, nil
}

func throwBall(HP int) bool {
	chance := 100 - HP
	if chance < 0 {
		chance = 0
	}
	return rand.Intn(100) < chance
}

func commandCatch(cfg *Config, args []string) error {
	attemptPokemon, hp, err := getStats(cfg, args)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	name := attemptPokemon.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	time.Sleep(1 * time.Second)
	if throwBall(hp) {
		fmt.Printf("You caught %s!\n", name)
		cfg.Pokedex = append(cfg.Pokedex, Pokedex{pokemon: attemptPokemon})
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
