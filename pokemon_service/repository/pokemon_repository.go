package repository

import (
	service "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/service"
	model "github.com/PasHdez/ondemand-go-bootcamp/models"
)

// pokemonRepository is the implementation of PokemonRepository interface
type pokemonRepository struct {
	Service service.Service
}

// PokemonRepository is the interface that provides the pokemonRepository methods
type PokemonRepository interface {
	GetPokemons() ([]model.Pokemon, error)
	GetPokemon(id int) (model.Pokemon, error)
	GetPokemonsPars(pType string, items int, itemsPerWorker int) ([]model.Pokemon, error)
}

// NewPokemonRepository returns a new instance of PokemonRepository
func NewPokemonRepository(s service.Service) PokemonRepository {
	return &pokemonRepository{s}
}

// GetPokemons returns all pokemons
func (p *pokemonRepository) GetPokemons() ([]model.Pokemon, error) {
	pokemons, err := p.Service.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// GetPokemon returns a pokemon by id
func (p *pokemonRepository) GetPokemon(id int) (model.Pokemon, error) {
	pokemon, err := p.Service.GetById(id)
	if err != nil {
		return model.Pokemon{}, err
	}
	return pokemon, nil
}

// GetPokemonsPars returns a list of pokemons by type and quantity
func (p *pokemonRepository) GetPokemonsPars(pType string, items int, itemsPerWorker int) ([]model.Pokemon, error) {
	pokemons, err := p.Service.GetByPars(pType, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}
