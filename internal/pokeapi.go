package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type PageConfig struct {
	next string
	prev string
}

type apiResponse struct {
	count int
	next string
	previous string
	results []locationResult
}

type locationResult struct {
	name string
	url string
}

func request(url string) (*apiResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Request failed with status %d", res.StatusCode))
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