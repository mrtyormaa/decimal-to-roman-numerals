package main

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

func TestServerStarts(t *testing.T) {
	go func() {
		// Start the server in a separate goroutine
		r := api.InitRouter()
		if err := r.Run(":8001"); err != nil {
			t.Errorf("failed to start server: %v", err)
		}
	}()

	// Give some time for the server to start
	time.Sleep(100 * time.Millisecond)

	// Send a test request to the server
	resp, err := http.Get("http://localhost:8001/")
	assert.NoError(t, err, "Failed to send request to server")
	defer resp.Body.Close()

	// Check if the server responds with a success status code
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Server did not respond with expected status code")
}

func TestGetPort(t *testing.T) {
	// Save the original PORT value and defer restoration
	originalPort := os.Getenv("PORT")
	defer os.Setenv("PORT", originalPort)

	tests := []struct {
		envPort      string
		expectedPort int
	}{
		{"8080", 8080},    // Valid port number
		{"invalid", 8001}, // Invalid port number
		{"", 8001},        // No port set, should use default
	}

	for _, test := range tests {
		// Set the PORT environment variable
		os.Setenv("PORT", test.envPort)

		// Get the port using the function
		port := getPort()

		// Check if the returned port matches the expected port
		if port != test.expectedPort {
			t.Errorf("Expected port %d, but got %d", test.expectedPort, port)
		}
	}
}
