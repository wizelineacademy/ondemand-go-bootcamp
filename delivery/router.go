package delivery

import (
	"lolidelgado/github-users/controller"
	"net/http"
)

func Setup(githubUser *controller.GithubUser) {
	http.HandleFunc("/github-users", githubUser.GithubUsersHandler)
}
