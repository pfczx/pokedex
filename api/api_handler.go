package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pfczx/pokedex/iternal/cashing"
)

var cache = cashing.NewCache(1 * time.Minute)

type Location struct {
	Name string `json:"name"`
}
type Locations struct {
	List []Location `json:"results"`
}

func HandleMap(url string) ([]Location, error) {

	if data, exists := cache.Get(url); exists {
		res, err := handleCasheGetMap(data)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	res, errHttp := http.Get(url)
	if errHttp != nil {
		return nil, errHttp
	}
	defer res.Body.Close()

	dc := json.NewDecoder(res.Body)
	var locations Locations
	err := dc.Decode(&locations)
	//decoder to replace by io readers
	data, errormarshal := json.Marshal(locations)
	if errormarshal != nil {
		return nil, err
	}
	cache.Add(url, data)

	if err != nil {
		return nil, err
	}
	var result []Location
	for _, location := range locations.List {
		result = append(result, location)
	}
	return result, nil

}

func handleCasheGetMap(data []byte) ([]Location, error) {
	var locations Locations
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return nil, err
	}
	var result []Location
	for _, location := range locations.List {
		result = append(result, location)
	}
	return result, nil

}
