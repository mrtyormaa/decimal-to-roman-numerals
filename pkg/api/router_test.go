package api_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	router := api.InitRouter()

	t.Run("GET /", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "Please use the /api/v1 endpoint for API access.")
	})

	t.Run("GET /health", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "Decimal to Roman Numerals Converter")
	})

	t.Run("GET /api/v1/health", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/health", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "Decimal to Roman Numerals Converter")
	})

	t.Run("GET /api/v1/convert", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/convert?numbers=1", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		// Add assertions based on the expected response for this endpoint
	})

	t.Run("POST /api/v1/convert", func(t *testing.T) {
		payload := `{"ranges": [{"min": 1, "max": 3999}]}`
		req, _ := http.NewRequest("POST", "/api/v1/convert", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		// Add assertions based on the expected response for this endpoint
	})
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
		port := api.GetPort()

		// Check if the returned port matches the expected port
		if port != test.expectedPort {
			t.Errorf("Expected port %d, but got %d", test.expectedPort, port)
		}
	}
}
