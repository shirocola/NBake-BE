package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGenerateToken tests the GenerateToken function to ensure it returns a valid JWT
func TestGenerateToken(t *testing.T) {
	username := "testuser"
	token, err := GenerateToken(username)
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
		return
	}
	if token == "" {
		t.Errorf("Generated token is empty")
	}
}

// TestVerifyToken tests the VerifyToken function to ensure it properly validates a token
func TestVerifyToken(t *testing.T) {
	username := "testuser"
	token, _ := GenerateToken(username)

	// Test valid token
	retrievedUsername, err := VerifyToken(token)
	if err != nil {
		t.Errorf("VerifyToken() error = %v", err)
	}
	if retrievedUsername != username {
		t.Errorf("VerifyToken() got = %v, want = %v", retrievedUsername, username)
	}

	// Test invalid token
	_, err = VerifyToken("invalidtoken")
	if err == nil {
		t.Errorf("VerifyToken() error = %v, wantErr true", err)
	}
}

// TestRoleCheckMiddleware tests the RoleCheckMiddleware to ensure it blocks unauthorized roles
func TestRoleCheckMiddleware(t *testing.T) {
	// Setup
	requiredRole := "admin"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	testHandler := RoleCheckMiddleware(requiredRole)(handler)

	// Test with the correct role
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := context.WithValue(req.Context(), userRoleKey, "admin")
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test with the incorrect role
	req, _ = http.NewRequest("GET", "/", nil)
	ctx = context.WithValue(req.Context(), userRoleKey, "user")
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("Middleware did not block incorrect role: got %v want %v", status, http.StatusForbidden)
	}
}
