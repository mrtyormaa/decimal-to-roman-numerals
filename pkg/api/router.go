package api

import (
	docs "github.com/mrtyormaa/decimal-to-roman-numerals/docs"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api/roman"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
	}
	r.Use(middleware.Cors())

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	version := "/api/v1"
	docs.SwaggerInfo.BasePath = version

	v1 := r.Group(version)
	{
		v1.GET("/", roman.Healthcheck)
		v1.GET("/convert", roman.ConvertNumbersToRoman)
		v1.POST("/convert", roman.ConvertRangesToRoman)

	}

	return r
}
