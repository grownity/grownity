package db

import (
	"context"
	"github.com/google/go-github/v41/github"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	firebaseDb "firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type FbInterface interface {
	UpdateOrg(organization github.Organization) error
}

func FirebaseClient() (*firebaseDb.Client, error) {
	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: os.Getenv("DB_URL"),
	}
	json := []byte(os.Getenv("FB_SERVICE_ACCOUNT"))
	opt := option.WithCredentialsJSON(json)
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func (db *Database) UpdateOrg(organization *github.Organization) error {
	if err := db.fbClient.NewRef("organization/details").Set(context.Background(), organization); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
