package repository

import (
	"errors"
	"log"
	"strconv"

	service "pas.com/v1/infrastructure/service"
	model "pas.com/v1/models"
)

type pokemonRepository struct {
	Service service.Service
}

type PokemonRepository interface {
	GetPokemons() []model.Pokemon
	GetPokemon(id int) model.Pokemon
	GetPokemonsPars(pType string, items int, itemsPerWorker int) ([]model.Pokemon, error)
}

func NewPokemonRepository(s service.Service) PokemonRepository {
	return &pokemonRepository{s}
}

func convertToStruct(records [][]string) []model.Pokemon {
	var result []model.Pokemon
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		var tmp model.Pokemon = model.Pokemon{Id: id, Name: record[1], Url: record[2], Sprite: record[3]}

		result = append(result, tmp)
	}
	return result
}

func (p *pokemonRepository) GetPokemons() []model.Pokemon {
	records := p.Service.GetData()
	return convertToStruct(records)
}

func (p *pokemonRepository) GetPokemon(id int) model.Pokemon {
	records := p.Service.GetData()
	pokemons := convertToStruct(records)
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return pokemon
		}
	}
	return model.Pokemon{}
}

func (p *pokemonRepository) GetPokemonsOdd(items int, itemsPerWorker int) []model.Pokemon {
	records := p.Service.GetData()
	pokemons := convertToStruct(records)
	var result []model.Pokemon

	for _, pokemon := range pokemons {
		if pokemon.Id%2 != 0 {
			result = append(result, pokemon)
		}
	}
	return result
}

func (p *pokemonRepository) GetPokemonsEven(items int, itemsPerWorker int) []model.Pokemon {

	records := p.Service.GetData()
	pokemons := convertToStruct(records)

	var result []model.Pokemon

	for _, pokemon := range pokemons {
		if pokemon.Id%2 == 0 {
			result = append(result, pokemon)
		}
	}

	return result
}

func (p *pokemonRepository) GetPokemonsPars(pType string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	if pType == "odd" {
		return p.GetPokemonsOdd(items, itemsPerWorker), nil
	} else if pType == "even" {
		return p.GetPokemonsEven(items, itemsPerWorker), nil
	} else {
		return nil, errors.New("error: Invalid type ")
	}
}
