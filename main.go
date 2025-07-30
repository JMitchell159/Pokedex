package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

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
	result := temp[:idx + 1]
	return result;
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	cliMap := map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
	}
	cliMap["help"] = cliCommand{
		name:			"help",
		description:	"Displays a help message",
		callback:		help(cliMap),
	}

	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		valid := false
		clean := cleanInput(sc.Text())
		for command := range cliMap {
			if clean[0] == command {
				err := cliMap[command].callback()
				if err != nil {
					fmt.Errorf("%v", err)
				}
				valid = true
				break
			}
		}
		if !valid {
			fmt.Printf("Unknown command\n")
		}
	}
}