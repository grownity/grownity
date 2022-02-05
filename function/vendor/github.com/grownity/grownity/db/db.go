package db

import (
	firebaseDb "firebase.google.com/go/v4/db"
	"github.com/google/go-github/v41/github"
)

type Database struct {
	FbInterface
	fbClient *firebaseDb.Client
}

var client Database

func InitDB() {
	FirebaseClient()
	return
}

func GetClient() *Database {
	return &client
}
func (client *Database) UpdateOrganization(organization *github.Organization) error {
	return client.UpdateOrg(organization)
}
