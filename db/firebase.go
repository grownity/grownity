package db

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v41/github"

	firebase "firebase.google.com/go/v4"
	firebaseDb "firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type FbInterface interface {
	UpdateOrg(organization github.Organization) error
}

func FirebaseClient(url string, account string) (*firebaseDb.Client, error) {
	ctx := context.Background()
	databaseUrl := url
	var json_acc []byte
	if databaseUrl == "" {
		databaseUrl = os.Getenv("DB_URL")
	}
	if account == "" {
		json_acc = []byte(os.Getenv("FB_SERVICE_ACCOUNT"))
	} else {
		json_acc = []byte(account)
	}
	config := &firebase.Config{
		DatabaseURL: databaseUrl,
	}
	opt := option.WithCredentialsJSON(json_acc)
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
