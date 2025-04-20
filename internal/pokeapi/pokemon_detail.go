package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

func (c *Client) GetPokemonDetail(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if dat, exists := c.cache.Get(url); exists {
		var details Pokemon
		err := json.Unmarshal(dat, &details)
		if err != nil {
			return Pokemon{}, err
		}
		
		return details, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var details Pokemon
	err = json.Unmarshal(dat, &details)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return details, nil
}