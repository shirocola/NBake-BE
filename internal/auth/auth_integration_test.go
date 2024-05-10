package auth_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shirocola/NBake-BE/internal/auth"
)

func TestLoginIntegration(t *testing.T) {
	// Setup
	server := httptest.NewServer(auth.SetupRoutes()) // Assume SetupRoutes sets up your router
	defer server.Close()

	// Test data
	body := strings.NewReader(`{"username":"testuser", "password":"password"}`)
	req, _ := http.NewRequest("POST", server.URL+"/login", body)
	req.Header.Set("Content-Type", "application/json")

	// Execute
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("Failed to make request:", err)
	}

	// Verify
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", response.Status)
	}
	// Further checks can be made here, such as token validity, content checks, etc.
}
