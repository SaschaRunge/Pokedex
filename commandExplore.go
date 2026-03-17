package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) <= 1 {
		return errors.New("Missing argument for explore-command. Usage: explore {location}.")
	}
	location := args[1]
	fmt.Printf("Exploring %s... .\n", args[1])

	encounters, err := config.Client.GetEncounters(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range encounters {
		fmt.Printf("- %s\n", pokemon)
	}
	return nil
}
