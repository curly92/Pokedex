package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lower := strings.ToLower(trimmed)
	list_lower := strings.Split(lower, " ")
	return list_lower
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean_input := cleanInput(text)
		fmt.Println("Your command was: " + clean_input[0])
	}
}
