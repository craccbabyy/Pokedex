package main

import (
	"PokedexCLI/internal/pokeapi"
	"time"
)

var supportedCommands = make(map[string]cliCommand)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startREPL(cfg) // CLEAN THIS SHIT UP

}
