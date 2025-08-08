package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaAPIResponse, error) {
	url := fmt.Sprintf("%s/location-area", BaseURL)
	if pageURL != nil {
		url = *pageURL
	}

	if value, ok := c.cache.Get(url); ok {
		locations := LocationAreaAPIResponse{}
		err := json.Unmarshal(value, &locations)
		if err != nil {
			return LocationAreaAPIResponse{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationAreaAPIResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaAPIResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreaAPIResponse{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaAPIResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	var locations LocationAreaAPIResponse
	if err = json.Unmarshal(body, &locations); err != nil {
		return LocationAreaAPIResponse{}, fmt.Errorf("error unmarshaling data: %v", err)
	}
	c.cache.Add(url, body)
	return locations, nil
}

func (c *Client) ListLocationByName(name string) (LocationAreaByNameAPIResponse, error) {
	url := fmt.Sprintf("%s/location-area/%s", BaseURL, name)

	if value, ok := c.cache.Get(url); ok {
		location := LocationAreaByNameAPIResponse{}
		err := json.Unmarshal(value, &location)
		if err != nil {
			return LocationAreaByNameAPIResponse{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationAreaByNameAPIResponse{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaByNameAPIResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaByNameAPIResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	var pokemons LocationAreaByNameAPIResponse
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return LocationAreaByNameAPIResponse{}, err
	}

	c.cache.Add(url, body)
	return pokemons, nil
}
