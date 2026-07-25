package main

import (
	"fmt"
	//"path/filepath"
)

type location struct {
	Name string `json:"name"`
	// more later, I bet!
}

type response struct {
	Results  []location `json:"results"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
}

func init() { // func auto runs at start
	// var cfg config
	supportedCommands["map"] = cliCommand{
		name:   "map",
		desc:   "Display the next 20 locations in the Pokemon world",
		callbk: commandMap,

		// api endpoint : https://pokeapi.co/api/v2/location-area/{id or name}/
	}
}

func commandMap(cfg *config, args ...string) error { // scroll forward/next 20 locations
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.NextURL)
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

/*
	var urlString string
	if cfg.NextURL == "" {
		urlString = "https://pokeapi.co/api/v2/location-area/"
	} else {
		urlString = *cfg.NextURL
	}
	err := mapper(cfg, urlString)
	if err != nil {
		return err
	}
	return nil
}

func mapper(cfg *config, urlString string) error {
	res, err := http.Get(urlString)
	if err != nil {
		return err //error with request
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err // error reading bytes
	}
	data := response{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err //error unmarshaling
	}
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	*cfg.NextURL = data.Next
	*cfg.PreviousURL = data.Previous
	return nil
}
*/
