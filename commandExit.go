package main

import (
	"fmt"
	"os"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
)

func commandExit(config *pokeapi.Config, args ...string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}
