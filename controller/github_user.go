package controller

import (
	"encoding/json"
	"lolidelgado/github-users/usecase"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type GithubUser struct {
	render *render.Render
	g      *usecase.GithubUserUseCase
}

func NewGithubUser(r *render.Render, g *usecase.GithubUserUseCase) *GithubUser {
	return &GithubUser{r, g}
}

func (c *GithubUser) GetGithubUsers(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		githubUsers, err := c.g.FetchAll()
		if err != nil {
			panic(err)
		}
		json.NewEncoder(rw).Encode(githubUsers)
	}

}

func (c *GithubUser) GetGithubUserById(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	numericId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	githubUser, err := c.g.GetById(numericId)
	json.NewEncoder(rw).Encode(githubUser)
}
