package registry

import (
	"github.com/PasHdez/ondemand-go-bootcamp/interface/controller"
	ip "github.com/PasHdez/ondemand-go-bootcamp/interface/presenter"
	ir "github.com/PasHdez/ondemand-go-bootcamp/interface/repository"
	"github.com/PasHdez/ondemand-go-bootcamp/usecase/interactor"
	up "github.com/PasHdez/ondemand-go-bootcamp/usecase/presenter"
	ur "github.com/PasHdez/ondemand-go-bootcamp/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(r.db)
}

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
