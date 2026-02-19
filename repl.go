package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanInput := []string{}

	for _, substring := range strings.Split(text, " ") {
		if substring != "" {
			cleanInput = append(cleanInput, strings.ToLower(substring))
		}
	}

	if len(cleanInput) == 0 {
		return []string{""}
	}

	return cleanInput
}
