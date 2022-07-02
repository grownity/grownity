package firebase

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	conf "github.com/grownity/grownity/config"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Credentials struct {
	ProjectID     string `yaml:"project_id"`
	PrivateKeyID  string `yaml:"private_key_id"`
	PrivateKey    string `yaml:"private_key"`
	ClientEmail   string `yaml:"client_email"`
	ClientID      string `yaml:"client_id"`
	ClientCertURL string `yaml:"client_x509_cert_url"`
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
	c := conf.Get()
	fb := make(map[string]string)
	if c.DB.Provider == "Firebase" && c.DB.Endpoint == "" {
		for _, key := range fb_config_keys {
			fb[key] = os.Getenv("FB_" + strings.ToUpper(key))
		}
	} else {
		fb["project_id"] = c.DB.FirebaseCredentials.ProjectID
		fb["private_key_id"] = c.DB.FirebaseCredentials.PrivateKeyID
		fb["private_key"] = c.DB.FirebaseCredentials.PrivateKey
		fb["client_email"] = c.DB.FirebaseCredentials.ClientEmail
		fb["client_id"] = c.DB.FirebaseCredentials.ClientID
		fb["client_x509_cert_url"] = c.DB.FirebaseCredentials.ClientCertURL
	}
	fb["type"] = "service_account"
	fb["auth_uri"] = "https://accounts.google.com/o/oauth2/auth"
	fb["token_uri"] = "https://oauth2.googleapis.com/token"
	fb["auth_provider_x509_cert_url"] = "https://www.googleapis.com/oauth2/v1/certs"
	return fb
}

func FirebaseClient() (*firebase.App, error) {
	c := conf.Get()
	ctx := context.Background()
	json_acc, _ := json.Marshal(getFBConfig())
	databaseURL := c.DB.Endpoint
	if databaseURL == "" {
		databaseURL = os.Getenv("DB_URL")
	}
	config := &firebase.Config{
		DatabaseURL: databaseURL,
	}
	opt := option.WithCredentialsJSON(json_acc)
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return app, nil
}