package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pfczx/pokedex/iternal/caching"
)

var cache = cashing.NewCache(1 * time.Minute)
// map and mapb
type Location struct {
	Name string `json:"name"`
}
type Locations struct {
	List []Location `json:"results"`
}
//explore
type PokemonEncounter struct {
    Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Pokemon struct {
    Name string `json:"name"`
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

func HandleExplore(url string) ([]Pokemon,error){
	if data,exists := cache.Get(url);exists{
    res,err := handleCashGetExplore(data)
		if err !=nil{
			return nil,err
		}
		return res,nil
	}

	res,errHttp := http.Get(url)
	if errHttp  != nil{
		return nil,errHttp
	}
	var location LocationArea
	dc:= json.NewDecoder(res.Body)
	errDecoder := dc.Decode(&location)
	data,errormarshal := json.Marshal(location)
	if errormarshal != nil {
		return nil,errormarshal
	}
	cache.Add(url,data)
	if errDecoder != nil {
		return nil,errDecoder
	}
	var pokemons []Pokemon
	for _,pokemon := range location.PokemonEncounters{
		pokemons = append(pokemons, pokemon.Pokemon)
	}
	return pokemons,nil

	
} 
func handleCashGetExplore(data []byte) ([]Pokemon,error){
	var location LocationArea
	err := json.Unmarshal(data,&location)
	if err != nil{
		return nil,err
	}
	var pokemons []Pokemon
	for _,pokemon := range location.PokemonEncounters {
		pokemons = append(pokemons,pokemon.Pokemon)
	}
	return pokemons,nil
}







