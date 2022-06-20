package function

import (
	"fmt"
	"net/http"
	"os"

	"os/exec"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func GrownityOn(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	fmt.Fprintf(w, "$> pwd\n%s\n\n", wd)

	bytes, _ := exec.Command("ls", "-l").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l\n%s\n\n", bytes)
}
