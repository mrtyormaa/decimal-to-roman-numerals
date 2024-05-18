package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	router := api.InitRouter()

	t.Run("GET /api/v1/", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "Decimal to Roman Numerals Converter")
	})
}
