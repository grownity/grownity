package db

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v41/github"

	firebase "firebase.google.com/go/v4"
	firebaseDb "firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type FbInterface interface {
	UpdateOrg(organization github.Organization) error
}

var fb_config_keys []string = []string{
	"project_id",
	"private_key_id",
	"private_key",
	"client_email",
	"client_id",
	"client_x509_cert_url",
}

func getFBConfig() map[string]string {
	fb := make(map[string]string)
	for _, key := range fb_config_keys {
		fb[key] = os.Getenv("FB_" + strings.ToUpper(key))
	}
	fb["type"] = "service_account"
	fb["auth_uri"] = "https://accounts.google.com/o/oauth2/auth"
	fb["token_uri"] = "https://oauth2.googleapis.com/token"
	fb["auth_provider_x509_cert_url"] = "https://www.googleapis.com/oauth2/v1/certs"
	return fb
}

func FirebaseClient(url string, account string) (*firebaseDb.Client, error) {
	ctx := context.Background()
	databaseUrl := url
	var json_acc []byte
	if databaseUrl == "" {
		databaseUrl = os.Getenv("DB_URL")
	}
	if account == "" {
		json_acc, _ = json.Marshal(getFBConfig())
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
