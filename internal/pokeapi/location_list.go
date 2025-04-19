package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locations := Locations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return Locations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	var locations Locations
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, dat)
	return locations, nil
}
