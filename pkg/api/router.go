package api

import (
	"fmt"
	"os"
	"strconv"

	docs "github.com/mrtyormaa/decimal-to-roman-numerals/docs"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api/roman"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetPort() int {
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

// InitRouter initializes the Gin router with middleware, routes, and Swagger documentation.
func InitRouter() *gin.Engine {
	r := gin.Default()

	// Get global Monitor object and configure it
	m := middleware.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// Apply middleware to the router
	m.Use(r)
	r.Use(gin.Logger())
	r.Use(middleware.Cors())

	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
	}

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Define API version and base path for Swagger
	version := "/api/v1"
	docs.SwaggerInfo.BasePath = version

	port := GetPort()

	// Root endpoint to display message to use API version
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Please use the %s endpoint for API access. Visit http://localhost:%d/swagger/index.html", version, port),
		})
	})

	// Healthcheck endpoint at the root level
	r.GET("/health", roman.Healthcheck)

	// Group v1 routes
	v1 := r.Group(version)
	{
		v1.GET("/health", roman.Healthcheck)
		v1.GET("/convert", roman.ConvertNumbersToRoman)
		v1.POST("/convert", roman.ConvertRangesToRoman)
	}

	return r
}
