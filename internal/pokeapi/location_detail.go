package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetail(location string) (LocationDetail, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		details := LocationDetail{}
		err := json.Unmarshal(val, &details)
		if err != nil {
			return LocationDetail{}, err
		}

		return details, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetail{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetail{}, err
	}

	var details LocationDetail
	err = json.Unmarshal(dat, &details)
	if err != nil {
		return LocationDetail{}, err
	}

	c.cache.Add(url, dat)
	return details, nil
}