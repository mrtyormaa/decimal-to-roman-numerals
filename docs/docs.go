// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Asutosh",
            "email": "asutosh.satapathy@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/convert": {
            "get": {
                "description": "Converts a comma-separated list of integers(within the range of 1 to 3999) into their corresponding Roman numeral representations.\nThe response provides a unique, ascending list of Roman numerals. Leading zeroes, leading '+' signs, and extra spaces are supported.\nFor example, /convert?numbers=1,1,2,2,2,3,3 will return results for 1, 2, 3.\nThis endpoint also supports pluralized query formats, such as /convert?numbers=1,2 or /convert?numbers=1\u0026numbers=2,3.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Convert Integers to Roman Numerals",
                "operationId": "convertNumbersToRoman",
                "parameters": [
                    {
                        "type": "string",
                        "example": "\"52\"; \"1,4,9\"; \"01,02\"; \"1,52,098,+437\"",
                        "description": "Single integer or Comma-separated list of integers to be converted",
                        "name": "numbers",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/types.RomanNumeralResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/types.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "This endpoint accepts a JSON request body with multiple ranges of numbers(within the range of 1 to 3999), converting each to its Roman numeral equivalent.\nBoth 'min' and 'max' values in the range are inclusive. For example, the range 1-3 will generate results for 1, 2, and 3.\nThe response provides a unique list of numbers in ascending order from all specified ranges, sorted in ascending order. For example, ranges 3-4 and 2-5 will return results for 2, 3, 4, and 5 only once.\nNote that leading zeroes and leading '+' signs are not supported due to JSON limitations. Query parameters are not accepted; the request must be sent as a JSON object.\n",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Convert Ranges of Numbers to Roman Numerals",
                "operationId": "convertRangesToRoman",
                "parameters": [
                    {
                        "description": "List of number ranges to be converted",
                        "name": "ranges",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.RangesPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.RomanNumeralResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid JSON Payload",
                        "schema": {
                            "$ref": "#/definitions/types.JsonErrorResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns the health status of the service along with a message.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check service health",
                "operationId": "healthCheck",
                "responses": {
                    "200": {
                        "description": "Service is healthy",
                        "schema": {
                            "$ref": "#/definitions/types.HealthResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "[ERR1002] invalid input: please provide valid integers within the supported range (1-3999)"
                },
                "invalid_numbers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['8888']"
                    ]
                }
            }
        },
        "types.HealthResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Decimal to Roman Numerals Converter"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "types.JsonErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "[ERR1005] invalid JSON: JSON must contain only the 'ranges' key, which should be an array of one or more objects with 'min' and 'max' values. 'min' and 'max' values must be within 1 to 3999, and 'min' should not be greater than 'max'. No other keys are allowed."
                }
            }
        },
        "types.NumberRange": {
            "type": "object",
            "required": [
                "max",
                "min"
            ],
            "properties": {
                "max": {
                    "description": "The maximum value of the range (inclusive).",
                    "type": "integer",
                    "example": 20
                },
                "min": {
                    "description": "The minimum value of the range (inclusive).",
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "types.RangesPayload": {
            "type": "object",
            "required": [
                "ranges"
            ],
            "properties": {
                "ranges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.NumberRange"
                    }
                }
            }
        },
        "types.RomanNumeral": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "integer",
                    "example": 100
                },
                "roman": {
                    "type": "string",
                    "example": "C"
                }
            }
        },
        "types.RomanNumeralResponse": {
            "type": "object",
            "properties": {
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.RomanNumeral"
                    }
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Roman Numeral Converter API",
	Description:      "This API takes a range of decimals and converts it to roman numerals",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
