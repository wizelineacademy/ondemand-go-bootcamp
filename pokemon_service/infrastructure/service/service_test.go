package service

import (
	"testing"

	mock "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
	model "github.com/PasHdez/ondemand-go-bootcamp/models"
	"github.com/stretchr/testify/assert"
)

// mockCSV is a mock of CSV data
var mockCSV = [][]string{
	{"1", "Bulbasaur", "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/1.svg", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"},
	{"2", "Ivysaur", "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/2.svg", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png"},
	{"3", "Venusaur", "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/3.svg", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/3.png"},
	{"4", "Charmander", "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/4.svg", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/4.png"},
}

// mockPokemon is a mock of Pokemon slice
var mockPokemon = []model.Pokemon{
	{Id: 1, Name: "Bulbasaur", Url: "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/1.svg", Sprite: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"},
	{Id: 2, Name: "Ivysaur", Url: "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/2.svg", Sprite: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png"},
	{Id: 3, Name: "Venusaur", Url: "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/3.svg", Sprite: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/3.png"},
	{Id: 4, Name: "Charmander", Url: "https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/4.svg", Sprite: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/4.png"},
}

// TestGetAll is a test for GetAll function
func TestGetAll(t *testing.T) {
	tests := []struct {
		name string
		want []model.Pokemon
		data [][]string
		err  error
	}{
		{name: "GetAll Success", want: mockPokemon, data: mockCSV, err: nil},
		{name: "Error on GetAll", want: nil, data: nil, err: ErrPokemonsNotFound},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			csvData := mock.MockStore{}
			csvData.On("ReadData").Return(test.data, nil)
			srv, err := NewService(csvData)
			assert.Nil(t, err)

			pokemons, err := srv.GetAll()
			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, test.want, pokemons)
			}
		})
	}
}

// TestGetByID is a test for GetByID function
func TestGetById(t *testing.T) {
	tests := []struct {
		name      string
		parameter int
		want      model.Pokemon
		data      [][]string
		err       error
	}{
		{name: "GetById 1 Success", parameter: 1, want: mockPokemon[0], data: mockCSV, err: nil},
		{name: "GetById 2 Success", parameter: 2, want: mockPokemon[1], data: mockCSV, err: nil},
		{name: "GetById 3 Success", parameter: 3, want: mockPokemon[2], data: mockCSV, err: nil},
		{name: "GetById 4 Success", parameter: 4, want: mockPokemon[3], data: mockCSV, err: nil},
		{name: "GetById error", parameter: 5, want: model.Pokemon{}, data: mockCSV, err: ErrPokemonNotFound},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			csvData := mock.MockStore{}
			csvData.On("ReadData").Return(test.data, nil)
			srv, err := NewService(csvData)
			assert.Nil(t, err)

			pokemon, err := srv.GetById(test.parameter)
			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, test.want, pokemon)
			}
		})
	}
}

// TestGetByName is a test for GetByName function
func TestGetByPars(t *testing.T) {
	tests := []struct {
		name     string
		parType  string
		parItems int
		parIPW   int
		want     []model.Pokemon
		data     [][]string
		err      error
	}{
		{name: "GetByPars even Success", parType: "even", parItems: 2, parIPW: 1, want: []model.Pokemon{mockPokemon[1], mockPokemon[3]}, data: mockCSV, err: nil},
		{name: "GetByPars odd Success", parType: "odd", parItems: 2, parIPW: 1, want: []model.Pokemon{mockPokemon[0], mockPokemon[2]}, data: mockCSV, err: nil},
		{name: "GetByPars invalid parameter", parType: "invalid", parItems: 2, parIPW: 1, want: nil, data: mockCSV, err: ErrPokemonInvalidParameterType},
		{name: "GetByPars invalid max items per worker", parType: "even", parItems: 2, parIPW: MaxItemsPerWorkerQuantity + 1, want: nil, data: mockCSV, err: ErrMaxItemsPerWorkersReached},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			csvData := mock.MockStore{}
			csvData.On("ReadData").Return(test.data, nil)
			srv, err := NewService(csvData)
			assert.Nil(t, err)

			pokemons, err := srv.GetByPars(test.parType, test.parItems, test.parIPW)
			if err != nil {
				assert.ErrorIs(t, err, test.err)
			} else {
				assert.Equal(t, test.want, pokemons)
			}
		})
	}
}
