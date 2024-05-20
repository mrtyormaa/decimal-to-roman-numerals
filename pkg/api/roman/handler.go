package roman

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/types"
)

var converter RomanConverter = &BasicRomanConverter{}

// @BasePath /

// Healthcheck handles the API request to check the service health.
// @Summary Check service health
// @Description Returns the health status of the service along with a message.
// @ID healthCheck
// @Accept json
// @Produce json
// @Success 200 {object} types.HealthResponse "Service is healthy"
// @Router /health [get]
func Healthcheck(g *gin.Context) {
	response := gin.H{
		"status":  "success",
		"message": "Decimal to Roman Numerals Converter",
	}
	g.JSON(http.StatusOK, response)
}

// ConvertNumbersToRoman handles the API request to convert numbers to Roman numerals.
// @Summary Convert Integers to Roman Numerals
// @Description Converts a comma-separated list of integers(within the range of 1 to 3999) into their corresponding Roman numeral representations.
// @Description The response provides a unique, ascending list of Roman numerals. Leading zeroes, leading '+' signs, and extra spaces are supported.
// @Description For example, /convert?numbers=1,1,2,2,2,3,3 will return results for 1, 2, 3.
// @Description This endpoint also supports pluralized query formats, such as /convert?numbers=1,2 or /convert?numbers=1&numbers=2,3.
// @ID convertNumbersToRoman
// @Accept json
// @Produce json
// @Param numbers query string true "Single integer or Comma-separated list of integers to be converted" example("52"; "1,4,9"; "01,02"; "1,52,098,+437")
// @Success 200 {object} types.RomanNumeralResponse "Successful response"
// @Failure 400 {object} types.ErrorResponse "Invalid input"
// @Router /convert [get]
func ConvertNumbersToRoman(c *gin.Context) {
	// Get all query parameters
	queryParams := c.Request.URL.Query()

	// Check if there are any query parameters other than 'numbers'
	for param := range queryParams {
		if param != "numbers" {
			c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeInvalidParam).Error()})
			return
		}
	}

	// Get the numbers parameters from the query string
	numbersParams := c.QueryArray("numbers")

	// Check if the numbers parameter is missing
	if len(numbersParams) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeMissingNumbersParam).Error()})
		return
	}

	// Parse and validate the number list
	numbers, invalidNumbers := ParseNumberList(numbersParams)

	// If there are any invalid numbers, return an error response
	if len(invalidNumbers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           NewAppError(CodeInvalidInput).Error(),
			"invalid_numbers": invalidNumbers,
		})
		return
	}

	// Convert the numbers to Roman numerals
	results := ConvertNumbersToRomanNumerals(numbers)

	// Return the results as a JSON response
	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ParseNumberList parses and validates an array of comma-separated list of numbers
func ParseNumberList(numbersParams []string) ([]int, []string) {
	var numbers []int
	var invalidNumbers []string

	// Iterate over each numbers parameter
	for _, numbersParam := range numbersParams {
		numberStrings := strings.Split(numbersParam, ",")
		for _, numberString := range numberStrings {
			// Trim spaces
			numberString = strings.TrimSpace(numberString)
			if numberString == "" {
				invalidNumbers = append(invalidNumbers, "")
				continue // Skip empty strings
			}
			number, err := strconv.Atoi(numberString)
			if err != nil || number < LowerLimit || number > UpperLimit {
				invalidNumbers = append(invalidNumbers, numberString)
			} else {
				numbers = append(numbers, number)
			}
		}
	}

	return numbers, invalidNumbers
}

// ConvertNumbersToRomanNumerals converts a list of unique numbers to their Roman numeral equivalents
func ConvertNumbersToRomanNumerals(numbers []int) []types.RomanNumeral {
	uniqueNumbers := make(map[int]struct{})
	for _, number := range numbers {
		uniqueNumbers[number] = struct{}{}
	}

	var results []types.RomanNumeral
	for number := range uniqueNumbers {
		// todo: need to handle the error thrown by Convert
		// we never reach to error state as number is always valid
		roman, _ := converter.Convert(number)
		results = append(results, types.RomanNumeral{
			Decimal: uint(number),
			Roman:   roman,
		})
	}

	// Sort the results by decimal value in ascending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Decimal < results[j].Decimal
	})

	return results
}

// Function to check for duplicate `ranges` keys
func hasDuplicateRangesKey(data string) error {
	if strings.Count(data, "ranges") > 1 {
		return NewAppError(CodeInvalidJSONDuplicateKeys)
	}
	return nil
}

// ConvertRangesToRoman handles the API request to convert ranges of numbers to Roman numerals.
// @Summary Convert Ranges of Numbers to Roman Numerals
// @Description This endpoint accepts a JSON request body with multiple ranges of numbers(within the range of 1 to 3999), converting each to its Roman numeral equivalent.
// @Description Both 'min' and 'max' values in the range are inclusive. For example, the range 1-3 will generate results for 1, 2, and 3.
// @Description The response provides a unique list of numbers in ascending order from all specified ranges, sorted in ascending order. For example, ranges 3-4 and 2-5 will return results for 2, 3, 4, and 5 only once.
// @Description Note that leading zeroes and leading '+' signs are not supported due to JSON limitations. Query parameters are not accepted; the request must be sent as a JSON object.
// @Description
// @ID convertRangesToRoman
// @Accept json
// @Produce json
// @Param ranges body types.RangesPayload true "List of number ranges to be converted" example({"ranges": [{"min": 50, "max": 52}, {"min": 10, "max": 12}]})
// @Success 200 {object} []types.RomanNumeralResponse
// @Failure 400 {object} types.JsonErrorResponse "Invalid JSON Payload"
// @Router /convert [post]
func ConvertRangesToRoman(c *gin.Context) {
	var payload map[string]interface{}

	// Read the raw request body
	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": NewAppError(CodeFailedReadBody).Error()})
		return
	}

	// Check for duplicate keys
	if err := hasDuplicateRangesKey(string(rawBody)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return error if we detect query parameters
	if len(c.Request.URL.Query()) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeQueryParamInPostRequest).Error()})
		return
	}

	// Unmarshal the raw body into a map
	if err := json.Unmarshal(rawBody, &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeInvalidRange).Error()})
		return
	}

	// Check if the payload contains exactly one key "ranges" and the value is an array
	rangesData, ok := payload["ranges"].([]interface{})
	if !ok || len(payload) != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeInvalidRange).Error()})
		return
	}

	// If "ranges" array is empty, return an empty result
	if len(rangesData) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeInvalidRange).Error()})
		return
	}

	// Parse the "ranges" key into RangesPayload struct
	var rangesPayload types.RangesPayload
	rangesDataJSON, _ := json.Marshal(rangesData)
	if err := json.Unmarshal(rangesDataJSON, &rangesPayload.Ranges); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": NewAppError(CodeInvalidRange).Error()})
		return
	}

	// Process the ranges to generate a list of numbers
	numbers, err := ProcessRanges(rangesPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the numbers to Roman numerals
	results := ConvertNumbersToRomanNumerals(numbers)

	// Return the results as a JSON response
	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ProcessRanges processes the ranges and generates a list of numbers
func ProcessRanges(payload types.RangesPayload) ([]int, error) {
	var numbers []int

	for _, r := range payload.Ranges {
		if r.Min < LowerLimit || r.Max > UpperLimit || r.Min > r.Max {
			return nil, NewAppError(CodeInvalidRange)
		}
		for i := r.Min; i <= r.Max; i++ {
			numbers = append(numbers, i)
		}
	}

	return numbers, nil
}
