package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

// @title           Roman Numeral Convertor API
// @version         1.0
// @description     This API takes a range of decimals and converts it to roman numerals

// @contact.name   Asutosh
// @contact.email  asutosh.satapathy@gmail.com

// @host      localhost:8001
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	r := api.InitRouter()

	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
