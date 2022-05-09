package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alexis-wizeline/ondemand-go-bootcamp/interface/controller"
	"github.com/alexis-wizeline/ondemand-go-bootcamp/interface/repository"
)

func NewRouter(api *echo.Echo) *echo.Echo {
	pc := controller.NewPokemonController(repository.NewPokemonRepository())
	api.GET("/", test)

	pokemons := api.Group("pokemons")
	pokemons.GET("*", pc.GetPokemons)
	pokemons.GET("/:id", pc.GetPokemonById)

	return api
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "this is a test")
}
