package main

import (
	"math/rand"
	"time"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
	"github.com/SaschaRunge/Pokedex/internal/pokedex"
)

type config struct {
	Next        string
	Previous    string
	HttpsClient *pokeapi.HttpsClient
	Pokedex     *pokedex.Pokedex
	random      *rand.Rand
}

func NewConfig(next, previous string, client *pokeapi.HttpsClient) config {
	return config{
		Next:        next,
		Previous:    previous,
		HttpsClient: client,
		Pokedex:     pokedex.NewPokedex(),
		random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (c *config) GetRandomProbability() float64 {
	return c.random.Float64()
}
