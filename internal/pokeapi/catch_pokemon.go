package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// GET Pokemon FROM Location -- follows "ListLocation" format
func (c *Client) CatchPokemon(pokemonName string) (Pokemon, error) {
	if pokemonName == "" {
		return Pokemon{}, errors.New("pokemon name cannot be empty")
	}
	url := baseURL + "/pokemon/" + pokemonName

	// first check the cache for an entry with matching url
	if val, ok := c.cache.Get(url); ok {

		catchResp := Pokemon{}
		err := json.Unmarshal(val, &catchResp) // if found, unmarshal it
		if err != nil {
			return Pokemon{}, err
		}
		return catchResp, nil
	}
	// if no matching entry found, make a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	//get data from response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	catchResp := Pokemon{}
	err = json.Unmarshal(data, &catchResp)
	if err != nil {
		return Pokemon{}, err
	}
	// cache data and return
	c.cache.Add(url, data)
	return catchResp, nil
}
