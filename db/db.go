package db

import (
	config "github.com/grownity/grownity/config"

	firebaseDb "firebase.google.com/go/v4/db"
	"github.com/google/go-github/v41/github"
)

type Database struct {
	FbInterface
	fbClient *firebaseDb.Client
}

var client Database

func InitDB() (err error) {
	c := config.Get()
	cl, err := FirebaseClient(c.DB.Endpoint, c.DB.FB_account)
	if err != nil {
		return err
	}
	client.fbClient = cl
	return nil
}

func GetClient() *Database {
	return &client
}
func (client *Database) UpdateOrganization(organization *github.Organization) error {
	return client.UpdateOrg(organization)
}
