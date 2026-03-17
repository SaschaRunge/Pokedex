package main

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	entries := config.Pokedex.GetEntries()
	if len(entries) == 0 {
		fmt.Println("Your Pokedex is currently empty.")
	} else {
		fmt.Printf("Your Pokedex (%d entries):\n", len(entries))
	}

	for _, entry := range entries {
		fmt.Printf("  - %s\n", entry)
	}
	return nil
}
