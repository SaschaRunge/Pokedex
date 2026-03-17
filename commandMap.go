package main

import (
	"encoding/json"
	"fmt"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
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

func commandMap(config *config, args ...string) error {
	fmt.Println("Scrolling forward... .")
	url := config.Next
	if url == "" {
		fmt.Printf("You are on the last page\n")
		return nil
	}

	if locationAreasJSON, err := getLocationAreas(url, config.Client); err == nil {
		updateConfig(config, locationAreasJSON)
		printLocationAreas(locationAreasJSON)
		return nil
	} else {
		return err
	}
}

func commandMapb(config *config, args ...string) error {
	fmt.Println("Scrolling backwards... .")
	url := config.Previous
	if url == "" {
		fmt.Printf("You are on the first page\n")
		return nil
	}

	if locationAreasJSON, err := getLocationAreas(url, config.Client); err == nil {
		updateConfig(config, locationAreasJSON)
		printLocationAreas(locationAreasJSON)
		return nil
	} else {
		return err
	}
}

func updateConfig(config *config, locationAreasJSON locationAreas) {
	config.Previous = locationAreasJSON.Previous
	config.Next = locationAreasJSON.Next
}

func getLocationAreas(url string, client *pokeapi.Client) (locationAreas, error) {
	dat, err := client.GetData(url)
	if err != nil {
		return locationAreas{}, nil
	}

	locationAreasJSON := locationAreas{}
	err = json.Unmarshal(dat, &locationAreasJSON)
	if err != nil {
		return locationAreas{}, err
	}
	return locationAreasJSON, nil
}

func printLocationAreas(locationAreasJSON locationAreas) {
	fmt.Println("Found locations:")
	for _, result := range locationAreasJSON.Results {
		fmt.Printf("- %s\n", result.Name)
	}
}
