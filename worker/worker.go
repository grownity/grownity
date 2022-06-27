package worker

import (
	"fmt"

	"github.com/grownity/grownity/config"
	db "github.com/grownity/grownity/db"
	gh "github.com/grownity/grownity/github"
)

func Worker(conf string) (err error) {
	err = initClient(conf)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	configuration := config.Get()
	err = updateGithub(configuration.GitHub.Organization)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	return
}

func updateGithub(org string) (err error) {
	orgGH := gh.GetOrganization(org)
	err = db.GetClient().UpdateOrganization(orgGH)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	repos := gh.GetRepos(org)
	err = db.GetClient().UpdateRepos(repos)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	return nil
}

func initClient(conf string) (err error) {
	err = config.Init(conf)
	if err != nil {
		return fmt.Errorf("Error initializing Configuration: %s", err.Error())
	}
	fmt.Println("Initializing Database Client")
	err = db.InitDB()
	if err != nil {
		return fmt.Errorf("Error initializing database : %s", err.Error())
	}
	fmt.Println("Initializing GitHub Client")
	gh.GitHubInit()

	return nil
}
