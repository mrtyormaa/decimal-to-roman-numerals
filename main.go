package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

// Get the port as set via environment variable
// If it has not been set, default to 8001
func getPort() int {
	// Define the default port
	defaultPort := 8001

	// Read the port from the environment variable
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(defaultPort)
	}

	// Convert the port string to an integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = defaultPort
	}
	return port
}

// @title           Roman Numeral Converter API
// @version         1.0
// @description     This API takes a range of decimals and converts it to roman numerals

// @contact.name   Asutosh
// @contact.email  asutosh.satapathy@gmail.com

// @host      localhost:8001
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	port := getPort()
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
