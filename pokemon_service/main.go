package main

import (
	"net/http"

	controller "github.com/PasHdez/ondemand-go-bootcamp/controllers"
	db "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
	service "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/service"
	repository "github.com/PasHdez/ondemand-go-bootcamp/repository"
	router "github.com/PasHdez/ondemand-go-bootcamp/server"
)

const path = "data.csv"

func main() {
	store := db.NewStore(path)
	s, err := service.NewService(store)
	if err != nil {
		panic(err)
	}
	repository := repository.NewPokemonRepository(s)
	controller := controller.NewPokemonController(repository)
	appRouter := router.NewRoute(controller)

	http.ListenAndServe(":8090", appRouter.Router())
}
