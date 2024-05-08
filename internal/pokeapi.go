package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PageConfig struct {
	Cur string
	Next string
	Prev string
}

type apiResponse struct {
	next string
	previous string
	results []locationResult
}

type locationResult struct {
	name string
}

func GetLocations(config *PageConfig, forwards bool) ([]string, *PageConfig, error) {
	url := ""
	if config.Next == "" && config.Prev == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else if forwards {
		if config.Next == "" {
			url = config.Cur
		} else {
			url = config.Next
		}
	} else {
		if config.Prev == "" {
			url = config.Cur
		} else {
			url = config.Prev
		}
	}

	response, err := request(url)
	if err != nil {
		return nil, nil, err
	}
	locations := []string{}
	for _, result := range response.results {
		locations = append(locations, result.name)
	}
	newConfig := &PageConfig{}
	newConfig.Cur = url
	newConfig.Prev = response.previous
	newConfig.Next = response.next
	return locations, newConfig, nil
}

func request(url string) (*apiResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("request failed with a status %d", res.StatusCode)
	}
	if err != nil {
		return nil, err
	}
	data := &apiResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}