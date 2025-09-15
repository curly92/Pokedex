package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/curly92/Pokedex/internal/pokeapi"
	"github.com/curly92/Pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
	Client   *pokeapi.Client
}

var registry = make(map[string]cliCommand)

func cleanInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lower := strings.ToLower(trimmed)
	list_lower := strings.Split(lower, " ")
	return list_lower
}

func init() {
	registry["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	registry["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	registry["map"] = cliCommand{
		name:        "map",
		description: "Gives the next 20 locations",
		callback:    commandMap,
	}
	registry["mapb"] = cliCommand{
		name:        "mapb",
		description: "Gives the previous 20 locations",
		callback:    commandMapBack,
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		Next:     nil,
		Previous: nil,
		Client: pokeapi.NewClient(
			pokecache.NewCache(10 * time.Second),
		),
	}

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
			err := val.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
