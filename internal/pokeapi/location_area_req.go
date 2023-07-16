package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Hit the cache")
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResponse, nil
	}

	fmt.Println("Missed the cache")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	//add to the cache
	c.cache.Add(fullURL, data)

	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LoactionArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Hit the cache")
		locationArea := LoactionArea{}
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LoactionArea{}, err
		}
		return locationArea, nil
	}

	fmt.Println("Missed the cache")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LoactionArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LoactionArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LoactionArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LoactionArea{}, err
	}

	locationArea := LoactionArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LoactionArea{}, err
	}

	//add to the cache
	c.cache.Add(fullURL, data)

	return locationArea, nil
}
