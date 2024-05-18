package docs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwaggerFilesGenerated(t *testing.T) {
	// Check if the Swagger JSON file exists
	_, err := os.Stat("docs.go")
	assert.NoError(t, err, "Swagger JSON file not found")

	// Check if the Swagger JSON file exists
	_, err = os.Stat("swagger.json")
	assert.NoError(t, err, "Swagger JSON file not found")

	// Check if the Swagger YAML file exists
	_, err = os.Stat("swagger.yaml")
	assert.NoError(t, err, "Swagger YAML file not found")
}
