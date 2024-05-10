package user

import (
	"reflect"
	"testing"
)

func TestLogin(t *testing.T) {
	cases := []struct {
		name     string
		user     User
		expected bool
	}{
		{"valid user", User{Username: "testuser", Password: "testpassword"}, true},
		{"invalid user", User{Username: "admin", Password: "admin1"}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if res := c.user.Login(c.user.Username, c.user.Password); res != c.expected {
				t.Errorf("Expected %v, got %v", c.expected, res)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	expectedUser := User{ID: "1", Username: "John"}
	result, err := AddUser("John")
	if err != nil {
		t.Error("Error should be nil")
	}
	if !reflect.DeepEqual(expectedUser, result) {
		t.Errorf("Expected: %v, got: %v", expectedUser, result)
	}
}
