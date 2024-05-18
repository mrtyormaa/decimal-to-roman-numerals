package models_test

import (
	"encoding/json"
	"testing"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestRomanNumeralJSONMarshalling(t *testing.T) {
	// Create a sample RomanNumeral instance
	romanNumeral := models.RomanNumeral{
		Decimal: 10,
		Roman:   "X",
	}

	// Marshal the RomanNumeral instance to JSON
	jsonData, err := json.Marshal(romanNumeral)
	assert.NoError(t, err, "Error marshalling RomanNumeral to JSON")

	// Unmarshal the JSON data back to a RomanNumeral instance
	var unmarshalledRomanNumeral models.RomanNumeral
	err = json.Unmarshal(jsonData, &unmarshalledRomanNumeral)
	assert.NoError(t, err, "Error unmarshalling JSON to RomanNumeral")

	// Ensure that the unmarshalled RomanNumeral instance matches the original
	assert.Equal(t, romanNumeral, unmarshalledRomanNumeral, "Unmarshalled RomanNumeral does not match original")
}
