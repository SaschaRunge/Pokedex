package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	commands := getCommands()

	commands["help"].callback()
	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input = cleanInput(scanner.Text())
		if command, ok := commands[input[0]]; ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Callback returned error: %v\n", err)
			}

		} else {
			fmt.Printf("Unknown command")
		}
	}

}
