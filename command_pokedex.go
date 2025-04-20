package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, arguments ...string) error {
	pokemonList := cfg.caughtPokemon
	if len(pokemonList) == 0 {
		return errors.New("no pokemon caught")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range pokemonList {
		fmt.Printf("  - %s \n", pokemon.Name)
	}
	
	return nil
}