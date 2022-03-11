package router

import (
	"net/http"

	controller "github.com/PasHdez/ondemand-go-bootcamp/controllers"
	"github.com/julienschmidt/httprouter"
)

// appRouter is the implementation of AppRouter interface
type appRouter struct {
	controller controller.Controller
}

// AppRouter is the interface that provides the appRouter methods
type AppRouter interface {
	Router() http.Handler
}

// NewAppRouter returns a new instance of AppRouter
func NewRoute(c controller.Controller) AppRouter {
	return &appRouter{c}
}

// Router returns a httprouter.Router with all routes configured
func (c *appRouter) Router() http.Handler {
	router := httprouter.New()
	router.GET("/Pokemons", c.controller.GetPokemons)
	router.GET("/Pokemons/:id", c.controller.GetPokemon)
	router.GET("/PokemonsPars/:pType/:items/:itemsPerWorker", c.controller.GetPokemonsPars)
	return router
}
