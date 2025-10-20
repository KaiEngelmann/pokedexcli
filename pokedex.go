package main

import "github.com/kaiengelmann/pokedexcli/internal/pokeapi"

var pokedex []Pokedex

type Pokedex struct {
	pokemon pokeapi.Pokemon
}
