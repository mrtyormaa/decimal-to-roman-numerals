package roman_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api/roman"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	// Create a Gin router
	router := gin.Default()
	router.GET("/healthcheck", roman.Healthcheck)

	// Create a request to send to the above route
	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedBody := `{"message":"Decimal to Roman Numerals Converter","status":"success"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestConvertNumbersToRoman(t *testing.T) {
	// Create a Gin router
	router := gin.Default()
	router.GET("/convert", roman.ConvertNumbersToRoman)

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
			name:             "ValidInput_MultipleUnique",
			queryParam:       "numbers=1,5,10,5,10,1",
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
			expectedResponse: `{"error":"The 'numbers' query parameter is required"}`,
		},
		{
			name:             "MissingQueryParam_OtherParam",
			queryParam:       "number=1,2,3",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error": "Only 'numbers' query parameter is allowed"}`,
		},
		{
			name:             "AscendingOrder",
			queryParam:       "numbers=100,50,10",
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"results":[{"number":10,"roman":"X"},{"number":50,"roman":"L"},{"number":100,"roman":"C"}]}`,
		},
		{
			name:             "MultipleQueryParam_Valid",
			queryParam:       "numbers=50,10&numbers=100",
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

// TestParseNumberList tests the ParseNumberList function
func TestParseNumberList(t *testing.T) {
	tests := []struct {
		name            string
		input           []string
		expectedNumbers []int
		expectedInvalid []string
	}{
		{
			name:            "Valid list of numbers",
			input:           []string{"1,2,3", "4", "5,6,7,8"},
			expectedNumbers: []int{1, 2, 3, 4, 5, 6, 7, 8},
			expectedInvalid: nil,
		},
		{
			name:            "Numbers out of range",
			input:           []string{"0, 4000, 5000"},
			expectedNumbers: nil,
			expectedInvalid: []string{"0", "4000", "5000"},
		},
		{
			name:            "Non-numeric strings",
			input:           []string{"a, b, c"},
			expectedNumbers: nil,
			expectedInvalid: []string{"a", "b", "c"},
		},
		{
			name:            "Empty strings",
			input:           []string{"", "1,,2", "  "},
			expectedNumbers: []int{1, 2},
			expectedInvalid: []string{"", "", ""},
		},
		{
			name:            "Mixed valid and invalid entries",
			input:           []string{"1, abc, 2", "4000, 3"},
			expectedNumbers: []int{1, 2, 3},
			expectedInvalid: []string{"abc", "4000"},
		},
		{
			name:            "Multiple input arrays",
			input:           []string{"1, 2", "3, 4", "5"},
			expectedNumbers: []int{1, 2, 3, 4, 5},
			expectedInvalid: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers, invalidNumbers := roman.ParseNumberList(tt.input)
			if !reflect.DeepEqual(numbers, tt.expectedNumbers) {
				t.Errorf("expected numbers %v, got %v", tt.expectedNumbers, numbers)
			}
			if !reflect.DeepEqual(invalidNumbers, tt.expectedInvalid) {
				t.Errorf("expected invalid numbers %v, got %v", tt.expectedInvalid, invalidNumbers)
			}
		})
	}
}

// TestConvertNumbersToRomanNumerals tests the ConvertNumbersToRomanNumerals function
func TestConvertNumbersToRomanNumerals(t *testing.T) {
	tests := []struct {
		input    []int
		expected []models.RomanNumeral
	}{
		{[]int{1, 2, 3}, []models.RomanNumeral{{Decimal: 1, Roman: "I"}, {Decimal: 2, Roman: "II"}, {Decimal: 3, Roman: "III"}}},
		{[]int{10, 20, 10}, []models.RomanNumeral{{Decimal: 10, Roman: "X"}, {Decimal: 20, Roman: "XX"}}},
		{[]int{}, []models.RomanNumeral{}},
		{[]int{3999}, []models.RomanNumeral{{Decimal: 3999, Roman: "MMMCMXCIX"}}},
	}

	for _, test := range tests {
		result := roman.ConvertNumbersToRomanNumerals(test.input)
		if !equalRomanNumeralSlices(result, test.expected) {
			t.Errorf("ConvertNumbersToRomanNumerals(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Helper functions to compare slices for testing

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalRomanNumeralSlices(a, b []models.RomanNumeral) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TestProcessRanges tests the ProcessRanges function
func TestProcessRanges(t *testing.T) {
	tests := []struct {
		name          string
		input         models.RangesPayload
		expected      []int
		expectedError string
	}{
		{
			name: "Valid Ranges",
			input: models.RangesPayload{
				Ranges: []models.NumberRange{
					{Min: 1, Max: 5},
					{Min: 10, Max: 15},
				},
			},
			expected:      []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14, 15},
			expectedError: "",
		},
		{
			name: "Invalid Range Min Greater Than Max",
			input: models.RangesPayload{
				Ranges: []models.NumberRange{
					{Min: 20, Max: 10},
				},
			},
			expected:      nil,
			expectedError: "invalid range. each range must be within 1 to 3999 and min should not be greater than max",
		},
		{
			name: "Range Out of Bounds",
			input: models.RangesPayload{
				Ranges: []models.NumberRange{
					{Min: 0, Max: 10},
				},
			},
			expected:      nil,
			expectedError: "invalid range. each range must be within 1 to 3999 and min should not be greater than max",
		},
		{
			name:          "Empty Ranges",
			input:         models.RangesPayload{},
			expected:      []int{},
			expectedError: "",
		},
	}

	for _, test := range tests {
		result, err := roman.ProcessRanges(test.input)
		if test.expectedError != "" {
			if err == nil || err.Error() != test.expectedError {
				t.Errorf("ProcessRanges(%v) error = %v; want %v", test.input, err, test.expectedError)
			}
		} else {
			if err != nil {
				t.Errorf("ProcessRanges(%v) unexpected error = %v", test.input, err)
			}
			if !equalIntSlices(result, test.expected) {
				t.Errorf("ProcessRanges(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

// TestConvertRangesToRoman tests the ConvertRangesToRoman handler function
func TestConvertRangesToRoman(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name          string
		input         models.RangesPayload
		expected      []models.RomanNumeral
		expectedError string
	}{
		{
			name: "Valid Ranges",
			input: models.RangesPayload{
				Ranges: []models.NumberRange{
					{Min: 10, Max: 12},
					{Min: 15, Max: 15},
				},
			},
			expected: []models.RomanNumeral{
				{Decimal: 10, Roman: "X"},
				{Decimal: 11, Roman: "XI"},
				{Decimal: 12, Roman: "XII"},
				{Decimal: 15, Roman: "XV"},
			},
			expectedError: "",
		},
		{
			name: "Invalid Range",
			input: models.RangesPayload{
				Ranges: []models.NumberRange{
					{Min: 4000, Max: 5000},
				},
			},
			expected:      nil,
			expectedError: "invalid range. each range must be within 1 to 3999 and min should not be greater than max",
		},
	}

	for _, test := range tests {
		body, _ := json.Marshal(test.input)
		req, _ := http.NewRequest("POST", "/convert", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		r := gin.Default()
		r.POST("/convert", roman.ConvertRangesToRoman)
		r.ServeHTTP(w, req)

		if test.expectedError != "" {
			if w.Code != http.StatusBadRequest {
				t.Errorf("Expected status %v; got %v", http.StatusBadRequest, w.Code)
			}
			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			if response["error"] != test.expectedError {
				t.Errorf("Expected error %v; got %v", test.expectedError, response["error"])
			}
		} else {
			if w.Code != http.StatusOK {
				t.Errorf("Expected status %v; got %v", http.StatusOK, w.Code)
			}
			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			var results []models.RomanNumeral
			for _, r := range response["results"].([]interface{}) {
				rMap := r.(map[string]interface{})
				results = append(results, models.RomanNumeral{
					Decimal: uint(rMap["number"].(float64)),
					Roman:   rMap["roman"].(string),
				})
			}
			if !equalRomanNumeralSlices(results, test.expected) {
				t.Errorf("Expected results %v; got %v", test.expected, results)
			}
		}
	}
}
