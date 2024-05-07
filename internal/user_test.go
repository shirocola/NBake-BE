package user

import (
	"testing"
)

func TestLogin(t *testing.T) {
	u := User{}
	username := "testuser"
	password := "testpassword"

	// Test case 1
	if success := u.Login(username, password); !success {
		t.Errorf("login was not successful for user %s", username)
	}

	// Test case 2
	if u.Login("admin", "admin1") != false {
		t.Error("Expected: false, got: true")
	}
}
