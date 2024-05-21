package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

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

	port := api.GetPort()
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
