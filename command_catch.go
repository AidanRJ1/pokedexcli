package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, arguments ...string) error {
	if len(arguments) == 0 {
		return errors.New("no arguments provided")
	}
	pokemonName := arguments[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonDetail(pokemonName)
	if err != nil {
		return err
	}

	catchRate := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s... \n", pokemon.Name)
	time.Sleep(3 * time.Second)
	if catchRate > 40 {
		fmt.Printf("%s escaped! \n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught! \n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
