package db

import (
	"fmt"

	firebase "firebase.google.com/go"
	config "github.com/grownity/grownity/config"
	fb "github.com/grownity/grownity/db/firebase"
)

type ClientInterface interface {
	FirebaseInterface
}

type Database struct {
	ClientInterface
	fbApp *firebase.App
}

var client Database

func InitDB() (err error) {
	c := config.Get()
	switch c.DB.Provider {
	case "Firebase":
		client.fbApp, err = fb.FirebaseClient()
	default:
		return fmt.Errorf("Database provider %s not supported", c.DB.Provider)
	}
	return
}

func GetClient() *Database {
	return &client
}
