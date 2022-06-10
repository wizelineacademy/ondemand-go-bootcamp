//Package usecase processes the information from repository, making any needed calculations, data transformation, etc.
//If it recieved data it's from controllers, coming from delivery layer
package usecase

import (
	"lolidelgado/github-users/models"
	"lolidelgado/github-users/repository"
)

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
