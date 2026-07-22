package pokeapi

// "PokedexCLI/internal/pokecache"
const (
	baseURL = "https://pokeapi.co/api/v2" // URL responds with "SHALLOW LOCATIONS" (as opposed to the deep response)
) // this is the LIST endpoint

// Response - shallow locations

type RespShallowLoc struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
