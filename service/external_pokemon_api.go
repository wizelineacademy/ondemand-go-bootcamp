package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
)

type ExternalPokemonAPI interface {
	GetPokemonFromAPI(id int) (*model.PokemonAPI, error)
}

type externalPokemonAPI struct {
	url string
}

func NewExternalPokemonAPI() ExternalPokemonAPI {
	return &externalPokemonAPI{
		url: "https://pokeapi.co/api/v2/pokemon",
	}
}

func (s externalPokemonAPI) GetPokemonFromAPI(id int) (*model.PokemonAPI, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d", s.url, id))
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, errors.New("The Pokemon does not exist")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var pokemonAPI *model.PokemonAPI
	json.Unmarshal(responseData, &pokemonAPI)
	return pokemonAPI, nil
}
