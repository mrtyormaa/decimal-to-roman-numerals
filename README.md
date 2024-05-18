# decimal-to-roman-numerals

[![codecov](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals/graph/badge.svg?token=WCPsoNnQEy)](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals)

## Overview

This repository provides an API endpoint to convert decimal number ranges to Roman numerals.

## Folder structure

```
decimal-to-roman-numerals/
|-- bin/
|-- cmd/
|   |-- server/
|       |-- main.go
|-- config/
|-- docs/
|   |-- docs.go
|   |-- swagger.json
|   |-- swagger.yaml
|-- pkg/
|   |-- api/
|       |-- roman/
|           |-- handler.go
|       |-- router.go
|   |-- middleware/
|   |-- models/
|       |-- roman.go
|-- tests/
|-- Dockerfile
|-- docker-compose.yml
|-- go.mod
|-- go.sum
|-- README.md
```

### Explanation of Directories and Files:

1. **`bin/`**: Contains the compiled binaries.

2. **`cmd/`**: Main applications for this project. The directory name for each application should match the name of the executable.

    - **`main.go`**: The entry point.

3. **`pkg/`**: Libraries and packages that are okay to be used by applications from other projects. 

    - **`api/`**: API logic.
        - **`handler.go`**: HTTP handlers.
        - **`router.go`**: Routes.
    - **`models/`**: Data models.
    - **`middleware/`**: Middleware for various functionalities.

4. **`tests/`**: Test cases.

## Getting Started

### Prerequisites

- Go 1.21+
- Docker
- Docker Compose
- Make

### Installation

1. Clone the repository

```bash
git clone https://github.com/mrtyormaa/decimal-to-roman-numerals
```

2. Navigate to the directory

```bash
cd decimal-to-roman-numerals
```

3. Build and run the Docker containers

For unix based systems:
```bash
make setup && make build && make up
```

For windows:
```bash
make setup; make build; make up
```

### Environment Variables

You can set the environment variables in the `.env` file. Important environment variables will be defined as we progress

## API Documentation

The API is documented using Swagger and can be accessed at:

```
http://localhost:8001/swagger/index.html
```

### Endpoints

#### 1. Convert Numbers to Roman Numerals

This endpoint converts a comma-separated list of numbers to their corresponding Roman numeral representations.

- **URL**: `/convert`
- **Method**: `GET`
- **Parameters**:
  - `numbers` (required): Comma-separated list of integers to be converted. Each number must be within the range 1 to 3999.
- **Example**: `/convert?numbers=10,50,100`

#### Response

- **Status Code**: `200 OK`
- **Body**: JSON object containing the results.
  - `results`: An array of objects containing the decimal number and its Roman numeral representation.
    - `number`: Decimal number.
    - `roman`: Roman numeral representation.

#### Example

Request:
```bash
GET /convert?numbers=10,50,100
```

Response:

```json
{
  "results": [
    {"number": 10, "roman": "X"},
    {"number": 50, "roman": "L"},
    {"number": 100, "roman": "C"}
  ]
}
```

#### 2. Convert Ranges of Numbers to Roman Numerals

This endpoint converts multiple ranges of numbers to their corresponding Roman numeral representations.

- **URL**: `/convert`
- **Method**: `POST`
- **Body**: JSON object containing an array of number ranges.
  - `ranges`: An array of objects specifying number ranges.
    - `min`: The minimum value of the range (inclusive).
    - `max`: The maximum value of the range (inclusive).
- **Example Request**:
  ```json
  {
    "ranges": [
      {"min": 10, "max": 20},
      {"min": 50, "max": 55}
    ]
  }
  ```

#### Response

- **Status Code**: `200 OK`
- **Body**: JSON object containing the results.
  - `results`: An array of objects containing the decimal number and its Roman numeral representation.
    - `number`: Decimal number.
    - `roman`: Roman numeral representation.

#### Example

Request:
```http
POST /convert
Content-Type: application/json

{
  "ranges": [
    {"min": 10, "max": 20},
    {"min": 50, "max": 55}
  ]
}
```

Response:

```json
{
  "results": [
    {"number": 10, "roman": "X"},
    {"number": 11, "roman": "XI"},
    {"number": 12, "roman": "XII"},
    {"number": 13, "roman": "XIII"},
    {"number": 14, "roman": "XIV"},
    {"number": 15, "roman": "XV"},
    {"number": 16, "roman": "XVI"},
    {"number": 17, "roman": "XVII"},
    {"number": 18, "roman": "XVIII"},
    {"number": 19, "roman": "XIX"},
    {"number": 20, "roman": "XX"},
    {"number": 50, "roman": "L"},
    {"number": 51, "roman": "LI"},
    {"number": 52, "roman": "LII"},
    {"number": 53, "roman": "LIII"},
    {"number": 54, "roman": "LIV"},
    {"number": 55, "roman": "LV"}
  ]
}
```

## TODOs

- Endpoint Improvements
- Possibly Authentication!
- Possible Rate Limiter!
- Possibly Redis DB support!
- config file & env file
- possibly prometheus integration for observability
