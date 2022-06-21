package function

import (
	"fmt"
	"net/http"

	"github.com/grownity/grownity/worker"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func GrownityOn(w http.ResponseWriter, r *http.Request) {
	result := worker.Worker()
	fmt.Fprintf(w, "\n%s\n\n", result)
}
