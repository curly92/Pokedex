package main

import (
	"fmt"
)

func commandHelp(conf *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, value := range registry {
		fmt.Println(value.name + ": " + value.description)
	}
	return nil
}
