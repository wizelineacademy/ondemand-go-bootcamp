package main

import (
	"net/http"

	controller "github.com/PasHdez/ondemand-go-bootcamp/controllers"
	db "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
	service "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/service"
	router "github.com/PasHdez/ondemand-go-bootcamp/server"
	repository "github.com/PasHdez/ondemand-go-bootcamp/services"
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
