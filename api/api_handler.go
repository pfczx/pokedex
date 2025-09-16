package main

import (
	//"encoding/json"
	//"fmt"
	"encoding/json"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
}
type Locations struct {
	List []Location `json:"results"`
}

func handleMap(url string) ([]Location, error) {
	res, errHttp := http.Get(url)
	if errHttp != nil {
		return nil, errHttp
	}
	defer res.Body.Close()

	dc := json.NewDecoder(res.Body)
	var locations Locations
	err := dc.Decode(&locations)
	if err != nil {
		return nil, err
	}
	var result []Location
	for _,location := range locations.List {
		result = append(result,location)
	}
	return result, nil

}
