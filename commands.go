package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args []string) error
}

var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays names of 20 location areas in Pokemon",
		callback:    commandMap,
	},
	"next": {
		name:        "next",
		description: "Displays next 20 location areas in Pokemon",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "displays previous 20 locations",
		callback:    commandMapb,
	},
	"previous": {
		name:        "map previous",
		description: "displays previous 20 locations",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "provides details of selected location",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "catch a pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "inspect pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "list pokemon in pokedex",
		callback:    commandPokedex,
	},
	"save": {
		name:        "save",
		description: "save pokedex",
		callback:    commandSave,
	},
	"load": {
		name:        "load",
		description: "loads previous game",
		callback:    commandLoad,
	},
}
