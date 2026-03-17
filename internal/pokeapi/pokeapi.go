package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SaschaRunge/Pokedex/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(180 * time.Second),
	}
}

func (c *Client) GetData(url string) ([]byte, error) {
	dat, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, err
		}
		if res.StatusCode == 404 {
			err_string := fmt.Sprintf("404: %s not found.\nThis is likely due to a typo in an argument provided.", url)
			return []byte{}, errors.New(err_string)
		}

		defer res.Body.Close()

		dat, err = io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, err
		}

		c.cache.Add(url, dat)
	}

	return dat, nil
}

func (c *Client) GetEncounters(location string) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	dat, err := c.GetData(url)

	if err != nil {
		return []string{}, err
	}

	encountersJSON := encounters{}
	err = json.Unmarshal(dat, &encountersJSON)
	if err != nil {
		return []string{}, err
	}

	encounters := []string{}

	for _, pokemon := range encountersJSON.Pokemon_encounters {

		encounters = append(encounters, pokemon.Pokemon.Name)
	}

	return encounters, nil
}

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	dat, err := c.GetData(url)

	if err != nil {
		return Pokemon{}, err
	}

	pokemonJSON := Pokemon{}
	err = json.Unmarshal(dat, &pokemonJSON)
	if err != nil {
		return Pokemon{}, err
	}

	//fmt.Printf("TEST %s", pokemonJSON.Name)
	return pokemonJSON, nil
}
