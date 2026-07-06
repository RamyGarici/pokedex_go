package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", name)
	return getPokemon(url, cfg)
}