package models

type GithubUser struct {
	Id          int
	Login       string
	Name        string
	Company     string
	Bio         string
	PublicRepos int
}

func NewGithubUser(id int, login, name, company, bio string, publicrepos int) GithubUser {
	return GithubUser{Id: id, Login: login, Name: name, Company: company, Bio: bio, PublicRepos: publicrepos}
}
