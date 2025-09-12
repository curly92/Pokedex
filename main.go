package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var registry = make(map[string]cliCommand)


func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lower := strings.ToLower(trimmed)
	list_lower := strings.Split(lower, " ")
	return list_lower
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range registry {
		fmt.Println(value.name + ": " + value.description)
	}
	return nil
}

func init(){
registry["exit"] = cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}
	registry["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		clean_input := cleanInput(text)
		command := clean_input[0]
		val, ok := registry[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := val.callback()
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
