package models

type GithubUser struct {
	Id          int
	Login       string
	Name        string
	Company     string
	Bio         string
	PublicRepos int
}
