package main

import (
	"log"
	"lolidelgado/github-users/controller"
	"lolidelgado/github-users/delivery"
	"lolidelgado/github-users/repository"
	"lolidelgado/github-users/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func main() {
	const fileName = "github_users.csv"
	//repository
	githubUserRepo := repository.NewGithubUser(fileName)

	//useCase
	githubUserUseCase := usecase.NewGithubUser(githubUserRepo)

	//controllers for Rest
	httpRender := render.New()
	usersController := controller.NewGithubUser(httpRender, githubUserUseCase)

	httpRouter := mux.NewRouter()
	delivery.Setup(
		usersController,
		httpRouter,
	)

	//start server
	log.Fatal(http.ListenAndServe(":7000", httpRouter))
}
