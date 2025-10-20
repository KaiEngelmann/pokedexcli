package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kaiengelmann/pokedexcli/internal/pokeapi"
	"github.com/kaiengelmann/pokedexcli/internal/pokecache"
)

func main() {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	cfg := Config{
		PokeClient: pokeapi.NewClient(httpClient, "https://pokeapi.co/api/v2", pokecache.NewCache(5*time.Second)),
		Next:       nil,
		Previous:   nil,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		clean_input := cleanInput(text)
		cmd := clean_input[0]
		args := clean_input[1:]
		if command, ok := supportedCommands[cmd]; ok {
			err := command.callback(&cfg, args)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
