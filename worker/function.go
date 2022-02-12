package function

import (
	"fmt"
	db "github.com/grownity/grownity/db"
	git "github.com/grownity/grownity/github"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func GrownityOn(w http.ResponseWriter, r *http.Request) {
	db.InitDB()
	database := db.GetClient()
	database.UpdateOrg(git.GetOrganization())
	fmt.Fprint(w, "Hello, World!")
}
