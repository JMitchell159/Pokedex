package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	split := strings.Split(trimmed, " ")
	temp := make([]string, len(split))
	idx := -1
	for _, word := range split {
		if len(word) > 0 {
			idx++
			temp[idx] = strings.ToLower(word)
		}
	}
	result := temp[:idx+1]
	return result
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists the pokemon in a specified location area",
			callback:    explore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Has a chance to catch a pokemon, adding it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists the information on a specified pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all of the Pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

func startRepl(cfg *config) {
	sc := bufio.NewScanner(os.Stdin)
	cliMap := getCommands()
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		clean := cleanInput(sc.Text())
		if len(clean) == 0 {
			continue
		}
		command, ok := cliMap[clean[0]]
		if !ok {
			fmt.Printf("Unknown command\n")
			continue
		}
		if len(clean) > 1 && (clean[0] == "explore" || clean[0] == "catch" || clean[0] == "inspect") {
			err := command.callback(cfg, clean[1])
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		} else if clean[0] == "explore" || clean[0] == "catch" || clean[0] == "inspect" {
			fmt.Printf("must include another argument for %s\n", clean[0])
		} else {
			err := command.callback(cfg, "")
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}
}
