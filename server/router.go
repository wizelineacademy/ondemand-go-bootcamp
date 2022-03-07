package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	controller "pas.com/v1/controllers"
)

func Router() http.Handler {
	router := httprouter.New()
	router.GET("/Pokemons", controller.GetPokemons)
	router.GET("/Pokemons/:id", controller.GetPokemon)
	return router
}
