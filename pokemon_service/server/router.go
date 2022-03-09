package router

import (
	"net/http"

	controller "github.com/PasHdez/ondemand-go-bootcamp/controllers"
	"github.com/julienschmidt/httprouter"
)

type appRouter struct {
	controller controller.Controller
}

type AppRouter interface {
	Router() http.Handler
}

func NewRoute(c controller.Controller) AppRouter {
	return &appRouter{c}
}

func (c *appRouter) Router() http.Handler {
	router := httprouter.New()
	router.GET("/Pokemons", c.controller.GetPokemons)
	router.GET("/Pokemons/:id", c.controller.GetPokemon)
	router.GET("/PokemonsPars/:pType/:items/:itemsPerWorker", c.controller.GetPokemonsPars)
	return router
}
