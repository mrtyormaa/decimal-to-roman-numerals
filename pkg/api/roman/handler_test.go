package roman

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestConvertNumbersToRoman(t *testing.T) {
	// Create a Gin router
	router := gin.Default()
	router.GET("/convert", ConvertNumbersToRoman)

	// Test cases
	testCases := []struct {
		name             string
		queryParam       string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "ValidInput_Single",
			queryParam:       "numbers=10",
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"results":[{"number":10,"roman":"X"}]}`,
		},
		{
			name:             "ValidInput_Multiple",
			queryParam:       "numbers=1,5,10",
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"results":[{"number":1,"roman":"I"},{"number":5,"roman":"V"},{"number":10,"roman":"X"}]}`,
		},
		{
			name:             "InvalidInput_NonNumeric",
			queryParam:       "numbers=1,abc,10",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Invalid input. Please provide valid integers within the supported range (1-3999).","invalid_numbers":["abc"]}`,
		},
		{
			name:             "InvalidInput_OutOfRange",
			queryParam:       "numbers=5000,10000",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Invalid input. Please provide valid integers within the supported range (1-3999).","invalid_numbers":["5000","10000"]}`,
		},
		{
			name:             "InvalidInput_MixedOutOfRange",
			queryParam:       "numbers=1,3,32,5000,10000",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Invalid input. Please provide valid integers within the supported range (1-3999).","invalid_numbers":["5000","10000"]}`,
		},
		{
			name:             "MissingQueryParam_NoParam",
			queryParam:       "",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"The 'numbers' query parameter is required."}`,
		},
		{
			name:             "MissingQueryParam_OtherParam",
			queryParam:       "number=1,2,3",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"The 'numbers' query parameter is required."}`,
		},
		{
			name:             "AscendingOrder",
			queryParam:       "numbers=100,50,10",
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"results":[{"number":10,"roman":"X"},{"number":50,"roman":"L"},{"number":100,"roman":"C"}]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request with the given query parameter
			req, err := http.NewRequest("GET", "/convert?"+tc.queryParam, nil)
			assert.NoError(t, err)

			// Create a response recorder to record the response
			res := httptest.NewRecorder()

			// Serve the request using the Gin router
			router.ServeHTTP(res, req)

			// Check the response status code
			assert.Equal(t, tc.expectedStatus, res.Code)

			// Check the response body
			assert.JSONEq(t, tc.expectedResponse, strings.TrimSpace(res.Body.String()))

			// Additional check for ascending order
			if tc.expectedStatus == http.StatusOK {
				// Unmarshal response to extract results
				var response struct {
					Results []models.RomanNumeral `json:"results"`
				}
				assert.NoError(t, json.Unmarshal(res.Body.Bytes(), &response))

				// Check if the results are sorted in ascending order
				assert.True(t, sort.SliceIsSorted(response.Results, func(i, j int) bool {
					return response.Results[i].Decimal < response.Results[j].Decimal
				}))
			}
		})
	}
}
