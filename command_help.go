package main

import "fmt"

func commandHelp(cfg *config) error {
	cliMap := getCommands()
	fmt.Printf("Welcome to the Pokedex help menu!\nUsage\n\n")
	for _, val := range cliMap {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}
