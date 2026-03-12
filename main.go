package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/SaschaRunge/Pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	commands := getCommands()
	mapConfig := config{
		Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous: "",
	}

	cache := pokecache.NewCache(10 * time.Second)

	commands["help"].callback(&mapConfig, cache)
	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input = cleanInput(scanner.Text())
		if command, ok := commands[input[0]]; ok {
			err := command.callback(&mapConfig, cache)
			if err != nil {
				fmt.Printf("Callback returned error: %v\n", err)
			}

		} else {
			fmt.Printf("Unknown command \n")
		}
	}

}
