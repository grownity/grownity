package github

import (
	"github.com/google/go-github/v41/github"
)

var gitClient *github.Client

func GitHubInit() {
	gitClient = github.NewClient(nil)
}

/*
func GetRepos(org *github.Organization) *[]github.Repository {
	var repos []github.Repository

	repo, _, _ := gitClient.Repositories.Get(context.Background(), org.GetLogin())
}*/
