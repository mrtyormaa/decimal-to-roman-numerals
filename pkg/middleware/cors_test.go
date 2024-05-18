package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCorsMiddleware(t *testing.T) {
	// Create a new Gin engine
	router := gin.New()

	// Use the Cors middleware
	router.Use(middleware.Cors())

	// Define a test route
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route")
	})

	// Test cases with different origins
	tests := []struct {
		origin         string
		expectedOrigin string
	}{
		{"http://127.0.0.1", "http://127.0.0.1"},
		{"http://127.0.0.1:8001", "http://127.0.0.1:8001"},
		{"http://localhost", "http://localhost"},
		{"http://localhost:8001", "http://localhost:8001"},
	}

	for _, tc := range tests {
		t.Run(tc.origin, func(t *testing.T) {
			// Create a mock request to the test route
			req, err := http.NewRequest("GET", "/test", nil)
			assert.NoError(t, err, "Failed to create request")
			req.Header.Set("Origin", tc.origin)

			// Create a mock response recorder
			resp := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(resp, req)

			// Check if the response headers are set correctly
			assert.Equal(t, tc.expectedOrigin, resp.Header().Get("Access-Control-Allow-Origin"), "Access-Control-Allow-Origin header not set correctly for origin: "+tc.origin)
			assert.Equal(t, "true", resp.Header().Get("Access-Control-Allow-Credentials"), "Access-Control-Allow-Credentials header not set correctly for origin: "+tc.origin)
		})
	}
}
