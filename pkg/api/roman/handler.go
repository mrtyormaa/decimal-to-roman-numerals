package roman

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
)

// @BasePath /

// Healthcheck godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router / [get]
func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "ok")
}

// GetRoman godoc
// @Summary Get Roman Numeral
// @Description Get the roman numeral equivalent for a given decimal
// @Tags romans
// @Produce json
// @Success 200 {object} models.Roman "Successfully retrieved a Roman"
// @Router /GetRoman [get]
func GetRoman(c *gin.Context) {
	result := models.Roman{
		Decimal: 1,
		Roman:   "I",
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
