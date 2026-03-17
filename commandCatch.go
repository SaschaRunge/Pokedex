package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
)

func commandCatch(config *config, args ...string) error {
	if len(args) <= 1 {
		return errors.New("Missing argument for catch-command. Usage: catch <pokemon_name>.")
	}
	arg := args[1]
	fmt.Printf("Throwing a Pokeball at %s... .\n", args[1])

	pokemon, err := config.Client.GetPokemon(arg)
	if err != nil {
		return err
	}

	fmt.Println("...")
	if catch(config.GetRandomProbability(), pokemon) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.Pokedex.Add(pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func catch(randomProbability float64, pokemon pokeapi.Pokemon) bool {
	k := 10. //steepness of curve
	x := float64(pokemon.BaseExperience) / 1000
	x0 := 200. / 1000 // 50% chance to catch at 150 base experience

	threshold := 1 / (1 + math.Exp(-k*(x-x0)))

	//fmt.Printf("base_experience: %v\n", pokemon.BaseExperience)
	//fmt.Printf("chance_to_catch: %f\n", 1-threshold)
	if randomProbability > threshold {
		return true
	}
	return false
}
