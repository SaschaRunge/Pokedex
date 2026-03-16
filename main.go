package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SaschaRunge/Pokedex/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	commands := getCommands()
	client := pokeapi.NewClient()
	mapConfig := pokeapi.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous: "",
		Client:   client,
	}

	commands["help"].callback(&mapConfig)
	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input = cleanInput(scanner.Text())
		if command, ok := commands[input[0]]; ok {
			err := command.callback(&mapConfig)
			if err != nil {
				fmt.Printf("Callback returned error: %v\n", err)
			}

		} else {
			fmt.Printf("Unknown command \n")
		}
	}

}
