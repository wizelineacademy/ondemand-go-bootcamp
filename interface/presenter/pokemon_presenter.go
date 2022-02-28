package presenter

import "github.com/PasHdez/ondemand-go-bootcamp/domain/model"

type pokemonPresenter struct {
}

type PokemonPresenter interface {
	ResponsePokemon(p []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemon(ps []*model.Pokemon) []*model.Pokemon {
	for _, p := range ps {
		p.Name = "Pokemon" + p.Name
	}
	return ps
}
