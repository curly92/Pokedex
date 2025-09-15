package main

import (
	"fmt"
)

func commandMapBack(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if conf.Previous == nil {
		fmt.Println("you're at the beginning")
		return nil
	}
	if conf.Previous != nil {
		url = *conf.Previous
	}
	result, err := conf.Client.GetLocations(url)
	if err != nil {
		return err
	}
	conf.Next = result.Next
	conf.Previous = result.Previous

	for _, r := range result.Results {
		fmt.Println(r.Name)
	}
	return nil
}
