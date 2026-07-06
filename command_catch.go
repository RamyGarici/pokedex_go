package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonName := args[0]

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	body, ok := cfg.Cache.Get(url)
	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			return fmt.Errorf("could not find pokemon %s", pokemonName)
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, body)
	}

	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	const threshold = 50
	if rand.Intn(pokemon.BaseExperience) < threshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}