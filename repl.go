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
	callback    func(*config) error
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
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
		err := command.callback(cfg)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
