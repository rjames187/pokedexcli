package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type PageConfig struct {
	Cur string
	Next string
	Prev string
}

type ApiResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


var interval, _ = time.ParseDuration("5m")
var cache *pokecache.Cache = pokecache.NewCache(interval)

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
	for _, result := range response.Results {
		locations = append(locations, result.Name)
	}
	newConfig := &PageConfig{}
	newConfig.Cur = url
	newConfig.Prev = response.Previous
	newConfig.Next = response.Next
	return locations, newConfig, nil
}

func request(url string) (*ApiResponse, error) {
	body, present := cache.Get(url)
	if present {
		fmt.Println("cache accessed!")
	}
	if !present {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("request failed with a status %d", res.StatusCode)
		}
		if err != nil {
			return nil, err
		}
		cache.Add(url, body)
	}
	data := &ApiResponse{}
	err := json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}