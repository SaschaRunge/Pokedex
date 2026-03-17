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
	config := NewConfig(
		"https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		"",
		pokeapi.NewClient(),
	)

	commands["help"].callback(&config)
	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input = cleanInput(scanner.Text())
		if command, ok := commands[input[0]]; ok {
			err := command.callback(&config, input...)
			if err != nil {
				fmt.Printf("\nCallback returned error: %v\n", err)
			}

		} else {
			fmt.Printf("Unknown command \n")
		}
	}

}
