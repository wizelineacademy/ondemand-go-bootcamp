package main

import (
	"log"
	"lolidelgado/github-users/controller"
	"lolidelgado/github-users/delivery"
	"lolidelgado/github-users/repository"
	"lolidelgado/github-users/usecase"
	"net/http"

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

	delivery.Setup(
		usersController,
	)

	//start server
	log.Fatal(http.ListenAndServe(":7000", nil))
}
