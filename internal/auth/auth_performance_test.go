package auth_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/shirocola/NBake-BE/internal/auth"
	vegeta "github.com/tsenart/vegeta/lib"
)

func TestLoginPerformance(t *testing.T) {
	rate := vegeta.Rate{Freq: 100, Per: time.Second} // 100 requests per second
	duration := 4 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:3000/auth/login", // Update this line if necessary
		Body:   []byte(`{"username":"testuser", "password":"password"}`),
	})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile response time: %s\n", metrics.Latencies.P99)
	fmt.Printf("Requests [Total: %d, Rate: %.2f/s]\n", metrics.Requests, metrics.Rate)
	fmt.Printf("Success [Total: %.2f%%]\n", metrics.Success*100)
	fmt.Printf("Latencies [Mean: %s, 95th: %s, 99th: %s]\n", metrics.Latencies.Mean, metrics.Latencies.P95, metrics.Latencies.P99)
	fmt.Printf("Bytes In [Total: %d, Mean: %.2f]\n", metrics.BytesIn.Total, metrics.BytesIn.Mean)
	fmt.Printf("Bytes Out [Total: %d, Mean: %.2f]\n", metrics.BytesOut.Total, metrics.BytesOut.Mean)
	fmt.Printf("Errors [Total: %d]\n", len(metrics.Errors))

	fmt.Printf("99th percentile response time: %s\n", metrics.Latencies.P99)
	if metrics.Success < 0.95 {
		t.Errorf("Success rate too low: %.2f", metrics.Success)
	}
}

func TestLoginRequest(t *testing.T) {
	server := httptest.NewServer(auth.SetupRoutes()) // Ensure this is correctly setting up routes
	defer server.Close()

	// Constructing a simple login request
	jsonData := []byte(`{"username":"testuser", "password":"password"}`)
	req, err := http.NewRequest("POST", server.URL+"/login", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal("Creating request failed:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("Failed to send request:", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v, body: %v", resp.Status, resp.Body)
	}
}
