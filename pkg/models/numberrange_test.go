package models_test

import (
	"encoding/json"
	"testing"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestNumberRangeJSONMarshalling(t *testing.T) {
	// Create a sample NumberRange instance
	numberRange := models.NumberRange{
		Min: 10,
		Max: 20,
	}

	// Marshal the NumberRange instance to JSON
	jsonData, err := json.Marshal(numberRange)
	assert.NoError(t, err, "Error marshalling NumberRange to JSON")

	// Unmarshal the JSON data back to a NumberRange instance
	var unmarshalledNumberRange models.NumberRange
	err = json.Unmarshal(jsonData, &unmarshalledNumberRange)
	assert.NoError(t, err, "Error unmarshalling JSON to NumberRange")

	// Ensure that the unmarshalled NumberRange instance matches the original
	assert.Equal(t, numberRange, unmarshalledNumberRange, "Unmarshalled NumberRange does not match original")
}

func TestRangesPayloadJSONMarshalling(t *testing.T) {
	// Create a sample RangesPayload instance
	rangesPayload := models.RangesPayload{
		Ranges: []models.NumberRange{
			{Min: 10, Max: 20},
			{Min: 30, Max: 40},
		},
	}

	// Marshal the RangesPayload instance to JSON
	jsonData, err := json.Marshal(rangesPayload)
	assert.NoError(t, err, "Error marshalling RangesPayload to JSON")

	// Unmarshal the JSON data back to a RangesPayload instance
	var unmarshalledRangesPayload models.RangesPayload
	err = json.Unmarshal(jsonData, &unmarshalledRangesPayload)
	assert.NoError(t, err, "Error unmarshalling JSON to RangesPayload")

	// Ensure that the unmarshalled RangesPayload instance matches the original
	assert.Equal(t, rangesPayload, unmarshalledRangesPayload, "Unmarshalled RangesPayload does not match original")
}
