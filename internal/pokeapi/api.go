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
			return LocationAreaAPIResponse{}, nil
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

	return locations, nil
}
