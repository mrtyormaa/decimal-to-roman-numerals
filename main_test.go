package main_test

import (
	"net/http"
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
	resp, err := http.Get("http://localhost:8001/api/v1/")
	assert.NoError(t, err, "Failed to send request to server")
	defer resp.Body.Close()

	// Check if the server responds with a success status code
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Server did not respond with expected status code")
}
