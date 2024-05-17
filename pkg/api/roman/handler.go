package roman

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
)

var converter RomanConverter = &BasicRomanConverter{}

// @BasePath /

func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "Decimal to Roman Numeral Converter")
}

// convert godoc
// @Summary Get Roman Numeral
// @Description Get the roman numeral equivalent for a given decimal(s) in ascending order
// @Tags romans
// @Produce json
// @Success 200 {object} models.RomanNumeral "Successfully retrieved Roman Numerals"
// @Router /convert [get]
func ConvertNumbersToRoman(c *gin.Context) {

	numbersParam := c.Query("numbers")
	if numbersParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The 'numbers' query parameter is required."})
		return
	}

	numberStrings := strings.Split(numbersParam, ",")
	var results []models.RomanNumeral
	var invalidNumbers []string

	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(strings.TrimSpace(numberString))
		if err != nil || number < 1 || number > 3999 {
			invalidNumbers = append(invalidNumbers, numberString)
		} else {
			roman, _ := converter.Convert(number)
			results = append(results, models.RomanNumeral{
				Decimal: uint(number),
				Roman:   roman,
			})
		}
	}

	if len(invalidNumbers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Please provide valid integers within the supported range (1-3999).", "invalid_numbers": invalidNumbers})
		return
	}

	// Sort the results in ascending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Decimal < results[j].Decimal
	})

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ranges godoc
// @Summary Get Roman Numerals for Ranges of Numbers
// @Description Get the roman numeral equivalent for given ranges in ascending order
// @Tags romans
// @Produce json
// @Success 200 {object} models.RomanNumeral "Successfully retrieved Roman Numerals"
// @Router /ranges [get]
func ConvertRangesToRoman(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not yet implemented."})
}
