package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// GET Pokemon FROM Location -- follows "ListLocation" format
func (c *Client) ListPokemon(locationName string) (RespExploreLoc, error) {
	if locationName == "" {
		return RespExploreLoc{}, errors.New("location name cannot be empty")
	}
	url := baseURL + "/location-area/" + locationName // not RespShallowLoc.name

	// first check the cache for an entry with matching url
	if val, ok := c.cache.Get(url); ok {

		exploreResp := RespExploreLoc{}
		err := json.Unmarshal(val, &exploreResp) // if found, unmarshal it
		if err != nil {
			return RespExploreLoc{}, err
		}
		return exploreResp, nil
	}
	// if no matching entry found, make a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLoc{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLoc{}, err
	}
	defer resp.Body.Close()
	//get data from response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLoc{}, err
	}
	exploreResp := RespExploreLoc{}
	err = json.Unmarshal(data, &exploreResp)
	if err != nil {
		return RespExploreLoc{}, err
	}
	// cache data and return
	c.cache.Add(url, data)
	return exploreResp, nil
}

/*
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`

	} `json:"pokemon_encounters"`
*/
