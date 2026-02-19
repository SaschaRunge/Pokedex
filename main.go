package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}

	for true {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input = cleanInput(scanner.Text())
		//fmt.Printf("Your command was: %s\n", input[0])
	}
}
