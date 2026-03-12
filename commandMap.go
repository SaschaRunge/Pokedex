package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type locationArea struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func commandMap(config *config) error {
	url := config.Next
	if url == "" {
		fmt.Printf("You are on the last page\n")
		return nil
	}

	return printLocationAreas(url, config)
}

func commandMapb(config *config) error {
	url := config.Previous
	if url == "" {
		fmt.Printf("You are on the first page\n")
		return nil
	}

	return printLocationAreas(url, config)
}

func printLocationAreas(url string, config *config) error {
	locationAreasJSON, err := getLocationAreas(url, config)
	for _, result := range locationAreasJSON.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return err
}

func getLocationAreas(url string, config *config) (locationAreas, error) {
	dat, exists := config.cache.Get(url)
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return locationAreas{}, err
		}
		defer res.Body.Close()

		dat, err = io.ReadAll(res.Body)
		if err != nil {
			return locationAreas{}, err
		}

		config.cache.Add(url, dat)
		fmt.Printf("Added contents of '%v' to cache\n", url)
	} else {
		fmt.Printf("Red contents of '%v' from cache\n", url)
	}

	locationAreasJSON := locationAreas{}
	err := json.Unmarshal(dat, &locationAreasJSON)
	if err != nil {
		return locationAreas{}, err
	}

	config.Previous = locationAreasJSON.Previous
	config.Next = locationAreasJSON.Next

	return locationAreasJSON, nil
}
