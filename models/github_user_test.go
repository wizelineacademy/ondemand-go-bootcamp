package models

import "testing"

func TestNewGithubUser(t *testing.T) {
	tests := []struct {
		name       string
		githubuser GithubUser
		expected   GithubUser
	}{
		{
			name: "Create a new github user",
			githubuser: GithubUser{
				Id:          1,
				Login:       "madodela",
				Name:        "Loli Delgado",
				Company:     "Wizeline",
				Bio:         "Software Engineer",
				PublicRepos: 20,
			},
			expected: GithubUser{
				Id:          1,
				Login:       "madodela",
				Name:        "Loli Delgado",
				Company:     "Wizeline",
				Bio:         "Software Engineer",
				PublicRepos: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGithubUser(tt.githubuser.Id, tt.githubuser.Login, tt.githubuser.Name, tt.githubuser.Company, tt.githubuser.Bio, tt.githubuser.PublicRepos)

			if got != tt.expected {
				t.Errorf("NewGithuUser(%+v,%+v,%+v,%+v,%+v,%+v) = %+v; expected %+v", tt.githubuser.Id, tt.githubuser.Login, tt.githubuser.Name, tt.githubuser.Company, tt.githubuser.Bio, tt.githubuser.PublicRepos, got, tt.expected)
			}
		})
	}
}
