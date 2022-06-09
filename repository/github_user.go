package repository

import (
	"encoding/csv"
	"io"
	"log"
	"lolidelgado/github-users/models"
	"strconv"
)

//aqui va a tener las fuciones para regresar la informacion de los githubusers, como un crud si fuera necesario

type GithubUser struct {
	githubUserCsv *csv.Reader
}

type IGithubUserRepository interface {
	FetchAll() ([]models.GithubUser, error)
}

func NewGithubUser(githubUserCsv *csv.Reader) *GithubUser {
	return &GithubUser{
		githubUserCsv,
	}
}

func (g *GithubUser) FetchAll() ([]models.GithubUser, error) {
	lines, err := g.readCsv()
	if err != nil {
		return []models.GithubUser{}, err
	}
	return arrayToGithubUser(lines), nil
}

func (g *GithubUser) readCsv() ([][]string, error) {
	var lines [][]string
	for {
		record, err := g.githubUserCsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if _, err := strconv.Atoi(record[0]); err == nil {
			lines = append(lines, record)
		}
	}
	return lines, nil
}

func arrayToGithubUser(lines [][]string) []models.GithubUser {
	var githubUsers []models.GithubUser
	for _, line := range lines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		publicRepos, err := strconv.Atoi(line[5])
		if err != nil {
			publicRepos = 0
		}
		data := models.GithubUser{
			Id:          id,
			Login:       line[1],
			Name:        line[2],
			Company:     line[3],
			Bio:         line[4],
			PublicRepos: publicRepos,
		}
		githubUsers = append(githubUsers, data)
	}
	return githubUsers
}
