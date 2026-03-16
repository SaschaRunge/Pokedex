package main

import (
	"fmt"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
)

func commandHelp(config *pokeapi.Config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
