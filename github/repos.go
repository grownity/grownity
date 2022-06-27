package github

import (
	"context"
	"time"

	"github.com/google/go-github/v41/github"
)

type RepoStats struct {
	Stars       *int `json:"stargazers,omitempty"`
	Forks       *int `json:"forks,omitempty"`
	Subscribers *int `json:"subscribers,omitempty"`
	Watchers    *int `json:"watchers,omitempty"`
}

type Timestamp struct {
	time.Time
}

func GetRepos(organization string) []*github.Repository {
	repos, _, _ := gitClient.Repositories.ListByOrg(context.Background(), organization, nil)
	return repos
}
