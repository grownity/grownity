package worker

import (
	"fmt"

	"github.com/grownity/grownity/config"
)

func Worker(conf string) string {
	err := config.Init(conf)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	configuration := config.Get()
	fmt.Printf(configuration.GitHub.Organization)
	fmt.Printf("I am the worker")
	return configuration.GitHub.Organization
}
