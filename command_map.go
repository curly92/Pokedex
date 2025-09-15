package main

import (
	"fmt"
)

func commandMap(conf *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if conf.Next != nil {
		url = *conf.Next
	}
	result, err := conf.Client.GetLocations(url)
	fmt.Println(url)
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
