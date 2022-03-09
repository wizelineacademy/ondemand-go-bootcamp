package main

import (
	"net/http"

	controller "pas.com/v1/controllers"
	db "pas.com/v1/infrastructure/database"
	service "pas.com/v1/infrastructure/service"
	router "pas.com/v1/server"
	repository "pas.com/v1/services"
)

func main() {
	path := "data.csv"
	store := db.NewStore(&path)
	s := service.NewService(store)
	repository := repository.NewPokemonRepository(s)
	controller := controller.NewPokemonController(repository)
	appRouter := router.NewRoute(controller)

	http.ListenAndServe(":8090", appRouter.Router())
}
