package github

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetOrg(t *testing.T) {
	GitHubInit()
	orgName := "kiali"
	org := GetOrganization(orgName)
	if org != nil {
		t.Fatalf(`Error getting organization.`)
	}

}
