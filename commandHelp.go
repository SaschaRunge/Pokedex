package main

import (
	"fmt"

	"github.com/SaschaRunge/Pokedex/internal/pokecache"
)

func commandHelp(config *config, cache *pokecache.Cache) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
