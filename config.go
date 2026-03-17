package main

import (
	"math/rand"
	"time"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
	"github.com/SaschaRunge/Pokedex/internal/pokedex"
)

type config struct {
	Next     string
	Previous string
	Client   *pokeapi.Client
	Pokedex  *pokedex.Pokedex
	random   *rand.Rand
}

func NewConfig(next, previous string, client *pokeapi.Client) config {
	return config{
		Next:     next,
		Previous: previous,
		Client:   client,
		Pokedex:  pokedex.NewPokedex(),
		random:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *config) GetRandomProbability() float64 {
	return c.random.Float64()
}
