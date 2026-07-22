package main

import (
	"errors"
	"fmt"
)

func init() { // func auto runs at start
	supportedCommands["mapb"] = cliCommand{
		name:   "mapb",
		desc:   "Display the previous 20 locations in the Pokemon world",
		callbk: commandMapb,
	}
}

func commandMapb(cfg *config) error { // scroll back/prev 20 locations
	if cfg.PreviousURL == nil { // if on first page
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.PreviousURL)
	if err != nil {
		return err
	}
	// set new Next and Previous
	cfg.NextURL = locationsResp.Next
	cfg.PreviousURL = locationsResp.Prev

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
