package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays names of locations in the Pokemon world. Each call shows the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of locations in the Pokemon world. Each call shows the previous 20 locations",
			callback:    commandMapb,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Tries to catch the specified pokemon. Takes a pokemon name as second argument.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspects a previously caught pokemon. Takes a pokemon name as second argument.",
			callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location>",
			description: "Displays names of pokemon encountered at the specified region. Takes location name as second argument.",
			callback:    commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
