package roman_test

import (
	"testing"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api/roman"
)

// Test that the BasicRomanConverter type satisfies the RomanConverter interface
func TestRomanConverterInterface(t *testing.T) {
	var _ roman.RomanConverter = (*roman.BasicRomanConverter)(nil)
}

// Test the converter function
func TestBasicRomanConverter_Convert(t *testing.T) {
	converter := roman.BasicRomanConverter{}

	testCases := []struct {
		input         int
		expected      string
		expectedError error
	}{
		{input: 0, expected: "", expectedError: roman.NewAppError(roman.CodeOutOfBounds)},
		{input: 4000, expected: "", expectedError: roman.NewAppError(roman.CodeOutOfBounds)},
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
