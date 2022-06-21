package worker

import (
	"fmt"

	"github.com/grownity/grownity/config"
)

func Worker() string {
	config.Init("")
	configuration := config.Get()
	fmt.Printf(configuration.GitHub.Organization)
	fmt.Printf("I am the worker")
	return configuration.GitHub.Organization
}
