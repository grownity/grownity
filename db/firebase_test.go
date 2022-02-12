package db

import (
	"os"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestFirebaseClient(t *testing.T) {
	if os.Getenv("DB_URL") != "" && os.Getenv("FB_SERVICE_ACCOUNT") != "" {
		_, err := FirebaseClient("", "")
		if err != nil {
			t.Fatalf(`Error in FB Client with empty data. %v`, err)
		}
	}
}
