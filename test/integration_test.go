package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"unicode"
)

// Helper function to check the response status code
func checkStatus(t *testing.T, w *httptest.ResponseRecorder, expectedStatus int) {
	if status := w.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}
}

// Helper function to unmarshal the response and check the results
func checkResponse(t *testing.T, w *httptest.ResponseRecorder, number int, expected string) {
	var response struct {
		Results []struct {
			Number int    `json:"number"`
			Roman  string `json:"roman"`
		} `json:"results"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if len(response.Results) != 1 || response.Results[0].Number != number || response.Results[0].Roman != expected {
		t.Errorf("handler returned unexpected body: got %v want {number: %d, roman: %s}", w.Body.String(), number, expected)
	}
}

// Helper function to unmarshal the response and check the results for the POST /convert endpoint
func checkPostResponse(t *testing.T, w *httptest.ResponseRecorder, expectedResults []struct {
	Number int
	Roman  string
}) {
	var response struct {
		Results []struct {
			Number int    `json:"number"`
			Roman  string `json:"roman"`
		} `json:"results"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if len(response.Results) != len(expectedResults) {
		t.Errorf("handler returned unexpected number of results: got %v want %v", len(response.Results), len(expectedResults))
	}

	for i, result := range response.Results {
		if result.Number != expectedResults[i].Number || result.Roman != expectedResults[i].Roman {
			t.Errorf("handler returned unexpected result at index %d: got {number: %d, roman: %s} want {number: %d, roman: %s}", i, result.Number, result.Roman, expectedResults[i].Number, expectedResults[i].Roman)
		}
	}
}

// Test cases for valid inputs for GET /api/v1/convert
func TestConvertHandlerValid(t *testing.T) {
	router := SetupRouter()
	testCases := []struct {
		number   int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		t.Run("Valid_"+strconv.Itoa(tc.number), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+strconv.Itoa(tc.number))
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, tc.number, tc.expected)
		})
	}
}

// Converts an integer to Roman string
func intToRoman(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	roman := ""
	roman += thousands[num/1000]
	roman += hundreds[(num%1000)/100]
	roman += tens[(num%100)/10]
	roman += ones[num%10]

	return roman
}

// Generate and test all 3999 numbers with a different algorithm to verify the algorithm validity
func TestConvertHandlerValidAnotherAlgorithm(t *testing.T) {
	router := SetupRouter()

	for i := 1; i <= 3999; i++ {
		t.Run("Valid_"+strconv.Itoa(i), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+strconv.Itoa(i))
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, i, intToRoman(i))
		})
	}
}

func TestConvertHandlerValidSpecial(t *testing.T) {
	router := SetupRouter()
	testCases := []struct {
		params   string
		number   int
		expected string
	}{
		{"01", 1, "I"},
		{"   04   ", 4, "IV"},
		{"+9", 9, "IX"},
	}

	for _, tc := range testCases {
		t.Run("Valid_"+tc.params, func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+tc.params)
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, tc.number, tc.expected)
		})
	}
}

// Generates all unicodes except numbers
func generateAllUnicodeCharsExceptNumbers() []string {
	var unicodeChars []string
	for i := rune(0); i <= unicode.MaxRune; i++ {
		if unicode.IsGraphic(i) && !unicode.IsNumber(i) {
			unicodeChars = append(unicodeChars, string(i))
		}
	}
	return unicodeChars
}

// Test cases for invalid inputs for GET /api/v1/convert
func TestConvertHandlerInvalid(t *testing.T) {
	router := SetupRouter()
	testCases := []string{
		"abc", "-1", "4000", "+0", "-1", "%1", "/1", "//1", "\\1", "~1", "^1",
		"°1", "1+2",
	}
	testCases = append(testCases, generateAllUnicodeCharsExceptNumbers()...)

	for _, tc := range testCases {
		t.Run("Invalid_"+tc, func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+tc)
			checkStatus(t, w, http.StatusBadRequest)
		})
	}
}

// Test cases for edge cases for GET /api/v1/convert
func TestConvertHandlerEdgeCases(t *testing.T) {
	router := SetupRouter()
	testCases := []struct {
		number   int
		expected string
	}{
		{1, "I"},
		{3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		t.Run("EdgeCase_"+strconv.Itoa(tc.number), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+strconv.Itoa(tc.number))
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, tc.number, tc.expected)
		})
	}
}

// Performance test to check the handler under load for GET /api/v1/convert
func TestConvertHandlerPerformance(t *testing.T) {
	router := SetupRouter()
	for i := 0; i < 1000; i++ {
		t.Run("LoadTest_"+strconv.Itoa(i), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers=123")
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, 123, "CXXIII")
		})
	}
}

// Test cases for POST /convert endpoint with valid inputs
func TestConvertRangesHandlerValid(t *testing.T) {
	router := SetupRouter()

	testCases := getRangesValidTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)
			checkPostResponse(t, w, tc.expectedResult)
		})
	}
}

// Test cases for POST /convert endpoint with invalid inputs
func TestConvertRangesHandlerInvalid(t *testing.T) {
	router := SetupRouter()

	testCases := getRangesInvalidTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)
		})
	}
}

// Test cases for edge cases for POST /convert endpoint
func TestConvertRangesHandlerEdgeCases(t *testing.T) {
	router := SetupRouter()
	testCases := getRangesEdgeTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)
			checkPostResponse(t, w, tc.expectedResult)

		})
	}
}

func TestConvertRangesHandlerEdgeCaseMaxValidRange(t *testing.T) {
	router := SetupRouter()

	testCases := getRangesEdgeTestCaseeMaxValidRange()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)

			var response struct {
				Results []struct {
					Number int    `json:"number"`
					Roman  string `json:"roman"`
				} `json:"results"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("failed to unmarshal response: %v", err)
				return
			}

			if len(response.Results) != 3999 {
				t.Errorf("handler returned unexpected number of results: got %v want %v", len(response.Results), 3999)
			}
			if response.Results[0].Number != 1 || response.Results[0].Roman != "I" {
				t.Errorf("handler returned unexpected first result: got {number: %d, roman: %s} want {number: 1, roman: I}", response.Results[0].Number, response.Results[0].Roman)
			}
			if response.Results[3998].Number != 3999 || response.Results[3998].Roman != "MMMCMXCIX" {
				t.Errorf("handler returned unexpected last result: got {number: %d, roman: %s} want {number: 3999, roman: MMMCMXCIX}", response.Results[3998].Number, response.Results[3998].Roman)
			}

		})
	}
}
