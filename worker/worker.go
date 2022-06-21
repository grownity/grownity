package worker

import (
	"fmt"

	"github.com/grownity/grownity/config"
)

func Worker() string {
	err := config.Init("")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	configuration := config.Get()
	fmt.Printf(configuration.GitHub.Organization)
	fmt.Printf("I am the worker")
	return configuration.GitHub.Organization
}
