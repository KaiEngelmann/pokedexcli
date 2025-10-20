package main

import (
	"fmt"
)

func commandHelp(cfg *Config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Print("help: Displays a help message\n")
	fmt.Print("exit: Exit the Pokedex\n")
	return nil
}
