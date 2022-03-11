package main

import (
	"log"
	"net/http"

	controller "github.com/PasHdez/ondemand-go-bootcamp/controllers"
	db "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
	service "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/service"
	repository "github.com/PasHdez/ondemand-go-bootcamp/repository"
	router "github.com/PasHdez/ondemand-go-bootcamp/server"
)

// path is the path to the csv file
const path = "data.csv"

// port is the port to listen
const port = ":8090"

// main is the entry point of the application
func main() {
	store := db.NewStore(path)
	s, err := service.NewService(store)
	if err != nil {
		panic(err)
	}
	repository := repository.NewPokemonRepository(s)
	controller := controller.NewPokemonController(repository)
	appRouter := router.NewRoute(controller)
	log.Printf("Listening on port %s", port)

	log.Fatal(http.ListenAndServe(port, appRouter.Router()))
}
