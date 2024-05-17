package roman

import (
	"errors"
	"testing"
)

func TestBasicRomanConverter_Convert(t *testing.T) {
	converter := BasicRomanConverter{}

	testCases := []struct {
		input         int
		expected      string
		expectedError error
	}{
		{input: 0, expected: "", expectedError: errors.New("input out of bounds, must be between 1 and 3999")},
		{input: 4000, expected: "", expectedError: errors.New("input out of bounds, must be between 1 and 3999")},
		{input: 1, expected: "I", expectedError: nil},
		{input: 3, expected: "III", expectedError: nil},
		{input: 4, expected: "IV", expectedError: nil},
		{input: 9, expected: "IX", expectedError: nil},
		{input: 10, expected: "X", expectedError: nil},
		{input: 58, expected: "LVIII", expectedError: nil},
		{input: 1994, expected: "MCMXCIV", expectedError: nil},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result, err := converter.Convert(tc.input)
			if tc.expectedError != nil {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Input: %d, Expected error: %v, Got error: %v", tc.input, tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Input: %d, Unexpected error: %v", tc.input, err)
				}
				if result != tc.expected {
					t.Errorf("Input: %d, Expected: %s, Got: %s", tc.input, tc.expected, result)
				}
			}
		})
	}
}
