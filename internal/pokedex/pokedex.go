package pokedex

import (
	"fmt"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
)

type Pokedex struct {
	content map[string]pokeapi.Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		content: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.content[pokemon.Name] = pokemon
}

func (p *Pokedex) GetEntries() []string {
	entries := []string{}
	for _, pokemon := range p.content {
		entries = append(entries, pokemon.Name)
	}

	return entries
}

func (p *Pokedex) DebugPrint() {
	fmt.Println(p.GetEntries())
}
