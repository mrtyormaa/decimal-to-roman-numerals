package api

import (
	"github.com/mrtyormaa/decimal-to-roman-numerals/docs"
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
		r.Use(middleware.Xss())
	}
	r.Use(middleware.Cors())

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/")
	{
		v1.GET("/", roman.Healthcheck)
		v1.GET("/GetRoman", roman.GetRoman)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
