package firebase

import (
	"context"
	"fmt"

	gh "github.com/grownity/grownity/github"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	github "github.com/google/go-github/v41/github"
)

func UpdateOrganization(app *firebase.App, organization *github.Organization) error {
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	defer client.Close()
	if _, err = client.Collection("github").Doc("orgDetails").Set(context.Background(), organization); err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func UpdateRepos(app *firebase.App, repos []*github.Repository) error {
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	defer client.Close()
	reps := make(map[string]*github.Repository)
	repsStats := make(map[string]map[string]*gh.RepoStats)
	for _, re := range repos {
		reps[*re.Name] = re
		repsStats[*re.Name] = make(map[string]*gh.RepoStats)
		repsStats[*re.Name][re.UpdatedAt.String()] = &gh.RepoStats{Stars: re.StargazersCount, Forks: re.ForksCount}
	}
	if _, err = client.Collection("github").Doc("repos").Set(context.Background(), reps); err != nil {
		fmt.Printf(err.Error())
		return err
	}
	w, err := client.Collection("github").Doc("reposStats").Get(context.Background())
	if w == nil {
		fmt.Printf(err.Error())
		return err
	}
	if !w.Exists() {
		// Init Collection
		if _, err = client.Collection("github").Doc("reposStats").Set(context.Background(), repsStats); err != nil {
			fmt.Printf(err.Error())
			return err
		}
	} else {
		// We need to add the new element
		if _, err = client.Collection("github").Doc("reposStats").Set(context.Background(), repsStats, firestore.MergeAll); err != nil {
			fmt.Printf(err.Error())
			return err
		}
	}
	return nil
}
