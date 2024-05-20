package api_test

import (
	"net/http"
	"net/http/httptest"
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
