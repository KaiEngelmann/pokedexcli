package main

import (
	"github.com/kaiengelmann/pokedexcli/internal/pokeapi"
	"github.com/kaiengelmann/pokedexcli/internal/pokecache"
)

type Config struct {
	PokeClient *pokeapi.Client
	Pokedex    []Pokedex
	Next       *string
	Previous   *string
	PokeCache  *pokecache.Cache
}
