package main

import (
	"fmt"

	worker "github.com/grownity/grownity/worker"
)

func main() {
	err := worker.Worker("./configuration.yaml")

	if err != nil {
		fmt.Errorf(err.Error())
	}
}
