package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
)

//go:generate mockery --name ExternalPokemonAPI --filename external_pokemon_api.go --outpkg mocks --structname ExternalPokemonAPIMock --disable-version-string
type ExternalPokemonAPI interface {
	GetPokemonFromAPI(id int) (*model.PokemonAPI, *model.ErrorHandler)
}

type externalPokemonAPI struct {
	url string
}

func NewExternalPokemonAPI() ExternalPokemonAPI {
	return &externalPokemonAPI{
		url: "https://pokeapi.co/api/v2/pokemon",
	}
}

func (s externalPokemonAPI) GetPokemonFromAPI(id int) (*model.PokemonAPI, *model.ErrorHandler) {
	response, err := http.Get(fmt.Sprintf("%s/%d", s.url, id))
	if err != nil {
		return nil, model.NewPokemonAPIIsNotReached(err.Error())
	}
	if response.StatusCode == http.StatusNotFound {
		return nil, model.NewGetPokemonFromAPINotFoundError(id)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, model.NewUnmarshalResponseBodyExternalService(err.Error())
	}
	var pokemonAPI *model.PokemonAPI
	json.Unmarshal(responseData, &pokemonAPI)
	return pokemonAPI, nil
}
