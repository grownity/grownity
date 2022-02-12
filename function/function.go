package function

import (
	"fmt"
	"net/http"
	//git "github.com/grownity/grownity/github"
	//db "github.com/grownity/grownity/db"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func GrownityOn(w http.ResponseWriter, r *http.Request) {
	//	db.InitDB()
	//	database := db.GetClient()
	//	database.UpdateOrg(git.GetOrganization())
	fmt.Fprint(w, "Hello, World!")
}
