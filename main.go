package main

import (
	"fmt"
	"strings"
)

func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lower := strings.ToLower(trimmed)
	list_lower := strings.Split(lower, " ")
	return list_lower
}

func main() {
	fmt.Println("Hello, World!")
}
