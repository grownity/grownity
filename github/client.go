package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v41/github"
)

var gitClient *github.Client

func GitHubInit() {
	gitClient = github.NewClient(nil)
}

func GetOrganization() *github.Organization {
	org, _, _ := gitClient.Organizations.Get(context.Background(), "kiali")
	fmt.Printf("%+v", org)
	return org
}

/*
func GetRepos(org *github.Organization) *[]github.Repository {
	var repos []github.Repository

	repo, _, _ := gitClient.Repositories.Get(context.Background(), org.GetLogin())
}*/
