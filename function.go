package function

import (
	"fmt"
	"net/http"
	"os"

	config "github.com/grownity/grownity/config"
	db "github.com/grownity/grownity/db"
	git "github.com/grownity/grownity/github"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func GrownityOn(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	if wd[len(wd)-1:] != "/" {
		wd += "/"
	}
	c, err := config.LoadFromFile(wd + "configuration.yaml")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	config.Set(c)

	err = db.InitDB()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	database := db.GetClient()
	database.UpdateOrg(git.GetOrganization("kiali"))
	fmt.Fprint(w, "Hello, World!")
}
