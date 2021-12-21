package function

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hellos, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hellos, World!")
		return
	}
	fmt.Fprintf(w, "Hellos, %s!", html.EscapeString(d.Name))
}
