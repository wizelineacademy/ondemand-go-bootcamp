package usecase

import (
	"lolidelgado/github-users/models"
	"lolidelgado/github-users/repository"
)

//aqui se va a procesar la informacion que lea del repositorio, calculos, o transformacion de datos
//si recibe datos son desde la capa de delivery

type GithubUserUseCase struct {
	githubUserRepo *repository.GithubUser
}

func NewGithubUser(githubUserRepo *repository.GithubUser) *GithubUserUseCase {
	return &GithubUserUseCase{
		githubUserRepo,
	}
}

func (g *GithubUserUseCase) FetchAll() ([]models.GithubUser, error) {
	return g.githubUserRepo.FetchAll()
}
