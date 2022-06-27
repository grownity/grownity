package db

import (
	"github.com/google/go-github/v41/github"
	fb "github.com/grownity/grownity/db/firebase"
)

type FirebaseInterface interface {
	UpdateOrganization(organization *github.Organization) error
	UpdateRepos(repos []*github.Repository)
}

func (client *Database) UpdateOrganization(organization *github.Organization) error {
	return fb.UpdateOrganization(client.fbApp, organization)
}

func (client *Database) UpdateRepos(repos []*github.Repository) error {
	return fb.UpdateRepos(client.fbApp, repos)
}
