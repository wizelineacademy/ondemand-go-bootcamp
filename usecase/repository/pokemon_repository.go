package repository

import "github.com/alexis-wizeline/ondemand-go-bootcamp/domain/model"

type PokemonRepository interface {
	GetPokemons() ([]*model.Pokemon, error)
	GetPokemonById(id uint64) (*model.Pokemon, error)
}
