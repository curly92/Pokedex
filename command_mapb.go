package main

import "fmt"

func commandMapBack(conf *config) error {
	var locationResponse LocationResponse
	if conf.Previous == nil {
		res, err := request("https://pokeapi.co/api/v2/location-area/")
		if err != nil {
			return err
		}
		locationResponse = res
	} else {
		res, err := request(*conf.Previous)
		if err != nil {
			return err
		}
		locationResponse = res
	}
	conf.Next = locationResponse.Next
	conf.Previous = locationResponse.Previous

	for _, val := range locationResponse.Results {
		fmt.Println(val.Name)
	}

	return nil
}
