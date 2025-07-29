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
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		clean := cleanInput(sc.Text())
		fmt.Printf("Your command was: %s\n", clean[0])
	}
}