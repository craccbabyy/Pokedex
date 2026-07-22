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

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespShallowLoc{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLoc{}, err
		}

		return locationsResp, nil
	}

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

	c.cache.Add(url, dat)
	return locationsResp, nil
}
