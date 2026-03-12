package main

import (
	"fmt"
	"os"

	"github.com/SaschaRunge/Pokedex/internal/pokecache"
)

func commandExit(config *config, cache *pokecache.Cache) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}
