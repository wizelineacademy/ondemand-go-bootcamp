package presenter

import "github.com/PasHdez/ondemand-go-bootcamp/domain/model"

type PokemonPresenter interface {
	ResponsePokemon(p []*model.Pokemon) []*model.Pokemon
}
