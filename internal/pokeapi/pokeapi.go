package pokeapi

import (
	"io"
	"net/http"

	"github.com/SaschaRunge/Pokedex/internal/pokecache"
)

type Client struct {
	cache *pokecache.Cache
	config
}

type config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
}

func (c *Client) GetData(url string) ([]byte, error) {
	dat, ok := c.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, err
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
