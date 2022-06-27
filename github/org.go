package github

import (
	"context"

	"github.com/google/go-github/v41/github"
)

func GetOrganization(organization string) *github.Organization {
	org, _, _ := gitClient.Organizations.Get(context.Background(), organization)
	return org
}
