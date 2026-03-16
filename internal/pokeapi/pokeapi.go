package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/SaschaRunge/Pokedex/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
	Client   *Client
}

type Client struct {
	cache *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		cache: pokecache.NewCache(10 * time.Second),
	}
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
