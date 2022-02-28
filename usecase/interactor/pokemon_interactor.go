package interactor

import (
	"github.com/PasHdez/ondemand-go-bootcamp/domain/model"
	"github.com/PasHdez/ondemand-go-bootcamp/usecase/presenter"
	"github.com/PasHdez/ondemand-go-bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (ps *pokemonInteractor) Get(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := ps.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return ps.PokemonPresenter.ResponsePokemon(p), nil
}
