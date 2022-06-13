package controller

import (
	"fmt"
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
			c.render.Text(rw, http.StatusInternalServerError, fmt.Sprintf("This is not fine🔥\nAnd the reason is: %s", err))
			return
		}
		c.render.JSON(rw, http.StatusOK, githubUsers)
	}

}

func (c *GithubUser) GetGithubUserById(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	numericId, err := strconv.Atoi(id)
	if err != nil {
		c.render.Text(rw, http.StatusBadRequest, "invalid id")
		return
	}
	githubUser, err := c.g.GetById(numericId)
	if err != nil {
		c.render.Text(rw, http.StatusInternalServerError, fmt.Sprintf("This is not fine🔥\nAnd the reason is: %s", err))
		return
	}
	c.render.JSON(rw, http.StatusOK, githubUser)
}
