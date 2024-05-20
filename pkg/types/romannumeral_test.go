package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumeralResponse(t *testing.T) {
	expected := RomanNumeralResponse{
		Results: []RomanNumeral{
			{Decimal: 100, Roman: "C"},
			{Decimal: 101, Roman: "CI"},
		},
	}

	data, err := json.Marshal(expected)
	assert.NoError(t, err)

	var actual RomanNumeralResponse
	err = json.Unmarshal(data, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestErrorResponse(t *testing.T) {
	expected := ErrorResponse{
		Error:          "[ERR1002] invalid input: please provide valid integers within the supported range (1-3999)",
		InvalidNumbers: []string{"8888"},
	}

	data, err := json.Marshal(expected)
	assert.NoError(t, err)

	var actual ErrorResponse
	err = json.Unmarshal(data, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestJsonErrorResponse(t *testing.T) {
	expected := JsonErrorResponse{
		Error: "[ERR1005] invalid JSON: JSON must contain only the 'ranges' key, which should be an array of one or more objects with 'min' and 'max' values. 'min' and 'max' values must be within 1 to 3999, and 'min' should not be greater than 'max'. No other keys are allowed.",
	}

	data, err := json.Marshal(expected)
	assert.NoError(t, err)

	var actual JsonErrorResponse
	err = json.Unmarshal(data, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestHealthResponse(t *testing.T) {
	expected := HealthResponse{
		Status:  "success",
		Message: "Decimal to Roman Numerals Converter",
	}

	data, err := json.Marshal(expected)
	assert.NoError(t, err)

	var actual HealthResponse
	err = json.Unmarshal(data, &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
