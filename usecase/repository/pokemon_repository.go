package repository

import "github.com/PasHdez/ondemand-go-bootcamp/domain/model"

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}
