package main

import (
	"encoding/csv"
	"log"
	"lolidelgado/github-users/controller"
	"lolidelgado/github-users/delivery"
	"lolidelgado/github-users/repository"
	"lolidelgado/github-users/usecase"
	"net/http"
	"os"

	"github.com/unrolled/render"
)

func main() {
	//csv reader
	file, err := os.Open("static/github_users.csv")
	if err != nil {
		log.Fatal("Invalid file: ", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	//repository
	githubUserRepo := repository.NewGithubUser(csvReader)

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
