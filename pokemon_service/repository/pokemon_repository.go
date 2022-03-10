package repository

import (
	"errors"

	service "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/service"
	model "github.com/PasHdez/ondemand-go-bootcamp/models"
)

type pokemonRepository struct {
	Service service.Service
}

type PokemonRepository interface {
	GetPokemons() ([]model.Pokemon, error)
	GetPokemon(id int) (model.Pokemon, error)
	GetPokemonsPars(pType string, items int, itemsPerWorker int) ([]model.Pokemon, error)
}

func NewPokemonRepository(s service.Service) PokemonRepository {
	return &pokemonRepository{s}
}

func (p *pokemonRepository) GetPokemons() ([]model.Pokemon, error) {
	pokemons, err := p.Service.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (p *pokemonRepository) GetPokemon(id int) (model.Pokemon, error) {
	pokemon, err := p.Service.GetById(id)
	if err != nil {
		return model.Pokemon{}, err
	}
	return pokemon, nil
}

func (p *pokemonRepository) GetPokemonsOdd(items int, itemsPerWorker int) ([]model.Pokemon, error) {
	pokemons, err := p.Service.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.Pokemon

	for _, pokemon := range pokemons {
		if pokemon.Id%2 != 0 {
			result = append(result, pokemon)
		}
	}
	return result, nil
}

func (p *pokemonRepository) GetPokemonsEven(items int, itemsPerWorker int) ([]model.Pokemon, error) {

	pokemons, err := p.Service.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.Pokemon

	for _, pokemon := range pokemons {
		if pokemon.Id%2 == 0 {
			result = append(result, pokemon)
		}
	}

	return result, nil
}

func (p *pokemonRepository) GetPokemonsPars(pType string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	if pType == "odd" {
		return p.GetPokemonsOdd(items, itemsPerWorker)
	} else if pType == "even" {
		return p.GetPokemonsEven(items, itemsPerWorker)
	} else {
		return nil, errors.New("error: Invalid type ")
	}
}
