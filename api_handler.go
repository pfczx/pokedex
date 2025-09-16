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

type LocationWrapper struct {
	Location Location `json:"location"`
}

type Locations struct {
	List []LocationWrapper `json:"results"`
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
	for _,locationWrapper := range locations.List {
		result = append(result,locationWrapper.Location)
	}
	return result, nil

}
