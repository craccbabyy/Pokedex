package main

import (
	"errors"
	"fmt"
)

func init() {
	supportedCommands["explore"] = cliCommand{
		name:   "explore",
		desc:   "See a list of all Pokemon located in an Area",
		callbk: commandExplore,
	}
}

// LIST THE POKEMON in the given LOCATION
func commandExplore(cfg *config, args ...string) error { // passing the location as argument
	// first validate that only one area arg was supplied
	if len(args) != 1 {
		return errors.New("please provide a single location name")
	}
	locationName := args[0]

	// call the new method on the client
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", pokemonResp.Location.Name)
	fmt.Println("Found Pokemon: ")
	//loop over the returned 'encounter' records
	for _, monster := range pokemonResp.PokemonEncounters {
		fmt.Println(monster.Pokemon.Name)
	}
	return nil

	//print each nested pokemon name on a new line

}
