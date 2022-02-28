package router

import (
	"github.com/PasHdez/ondemand-go-bootcamp/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/users", func(context echo.Context) error {
	// 	return c.User.GetUsers(context)
	// })

	e.GET("/pokemons", func(context echo.Context) error {
		return c.Pokemon.GetPokemons(context)
	})

	return e
}
