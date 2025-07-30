package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func help(commandMap map[string]cliCommand) func() error {
	return func() error {
		fmt.Print("Welcome to the Pokedex!\nUsage\n\n")
		for _, val := range commandMap {
			fmt.Printf("%s: %s\n", val.name, val.description)
		}
		return nil
	}
}