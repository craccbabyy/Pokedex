package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLoc, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	// check the cache first
	if val, ok := c.cache.Get(url); ok { // if found in cache
		locationsResp := RespShallowLoc{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil { // unmarshal data
			return RespShallowLoc{}, err
		}
		return locationsResp, nil //return
	}
	// if not found in cache, make a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLoc{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLoc{}, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLoc{}, err
	}
	locationsResp := RespShallowLoc{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLoc{}, err
	}

	// add entry to cache and return
	c.cache.Add(url, dat)
	return locationsResp, nil
}
