package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arguments ...string) error {
	if len(arguments) == 0 {
		return errors.New("no arguments provided")
	}
	details, err := cfg.pokeapiClient.GetLocationDetail(arguments[0])
	if err != nil {
		return err
	} 

	encounters := details.PokemonEncounters
	
	fmt.Printf("Exploring %s... \n", arguments[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters {
		pokemon := encounter.Pokemon
		fmt.Printf("- %s \n", pokemon.Name)
	}

	return nil
}
