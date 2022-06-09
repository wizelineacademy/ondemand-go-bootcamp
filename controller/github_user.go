package controller

import (
	"encoding/json"
	"lolidelgado/github-users/usecase"
	"net/http"

	"github.com/unrolled/render"
)

type GithubUser struct {
	render *render.Render
	g      *usecase.GithubUserUseCase
}

func NewGithubUser(r *render.Render, g *usecase.GithubUserUseCase) *GithubUser {
	return &GithubUser{r, g}
}

func (c *GithubUser) GithubUsersHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		githubUsers, err := c.g.FetchAll()
		if err != nil {
			panic(err)
		}
		enc := json.NewEncoder(rw)
		enc.Encode(githubUsers)
	}

}
