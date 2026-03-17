package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
)

func commandInspect(config *config, args ...string) error {
	if len(args) <= 1 {
		return errors.New("Missing argument for inspect-command. Usage: inspect <pokemon>.")
	}
	pokemonName := args[1]

	if pokemon, exists := config.Pokedex.Get(pokemonName); !exists {
		fmt.Printf("%s has not been caught yet.\n", pokemonName)
		return nil
	} else {
		fmt.Printf("%s", formatPokemon(pokemon))
	}

	return nil
}

func formatPokemon(pokemon pokeapi.Pokemon) string {
	sb := strings.Builder{}

	fmt.Fprintf(&sb, "Name: %s\n", pokemon.Name)
	fmt.Fprintf(&sb, "Height: %d\n", pokemon.Height)
	fmt.Fprintf(&sb, "Weight: %d\n", pokemon.Weight)
	fmt.Fprintf(&sb, "Stats:\n")

	for _, stat := range pokemon.Stats0 {
		fmt.Fprintf(&sb, "  - %s: %d\n", stat.Stats0Stat.Name, stat.BaseStat)
	}

	fmt.Fprintf(&sb, "Types:\n")

	for _, pokemonType := range pokemon.Types {
		fmt.Fprintf(&sb, "  - %s\n", pokemonType.Type.Name)
	}

	return sb.String()
}
