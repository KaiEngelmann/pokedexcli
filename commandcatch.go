package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getStats(cfg *Config, args []string) (string, int, error) {
	if len(args) == 0 {
		return "", 1, fmt.Errorf("please provide a pokemon to catch")
	}

	pokeName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokeName)

	attemptPokemon, err := cfg.PokeClient.CheckPokemon(url)
	if err != nil {
		return "", 1, fmt.Errorf("failed to find Pokemon: %w", err)
	}

	fmt.Printf("Exploring area: %s (ID: %d)\n", attemptPokemon.Name, attemptPokemon.ID)

	hp := 0
	for _, s := range attemptPokemon.Stats {
		if s.Stat.Name == "hp" {
			hp = s.BaseStat
			break
		}
	}

	return attemptPokemon.Name, hp, nil
}

func throwBall(HP int) bool {
	chance := 100 - HP
	if chance < 0 {
		chance = 0
	}
	return rand.Intn(100) < chance
}

func commandCatch(cfg *Config, args []string) error {
	name, hp, err := getStats(cfg, args)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	time.Sleep(1 * time.Second)
	if throwBall(hp) {
		fmt.Printf("You caught %s!\n", name)
		attemptPokemon, _ := cfg.PokeClient.CheckPokemon(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
		pokedex = append(pokedex, Pokedex{pokemon: attemptPokemon})
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
