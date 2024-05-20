# decimal-to-roman-numerals

[![Test and coverage](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions/workflows/ci.yml/badge.svg)](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals/graph/badge.svg?token=WCPsoNnQEy)](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtyormaa/decimal-to-roman-numerals)](https://goreportcard.com/report/github.com/mrtyormaa/decimal-to-roman-numerals)

## Overview

This repository provides an API endpoint to convert decimal number ranges to Roman numerals.

**TL;DR to Run and See Results**

Execute `make up` command and visit the links below.

| Link                             | Description                                                |
|----------------------------------|------------------------------------------------------------|
| [http://localhost:8001/swagger/index.html](http://localhost:8001/swagger/index.html) | Link to the Swagger API documentation                  |
| [http://localhost:8001/metrics](http://localhost:8001/metrics) | Link to the metrics endpoint used for Prometheus                |
| [http://localhost:9090](http://localhost:9090) | Link to the Prometheus instance                        |
| [http://localhost:3000](http://localhost:3000) | Link to the Grafana instance, login: `admin`, password: `admin` |


## Folder structure

```
decimal-to-roman-numerals/
|-- bin/
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
|-- test/
|-- Dockerfile
|-- docker-compose.yml
|-- go.mod
|-- go.sum
|-- main.go
|-- README.md
```

### Explanation of Directories and Files:

1. **`bin/`**: Contains the compiled binaries.

2. **`main.go`**: The entry point.

3. **`pkg/`**: Libraries and packages that are okay to be used by applications from other projects. 

    - **`api/`**: API logic.
        - **`handler.go`**: HTTP handlers.
        - **`router.go`**: Routes.
    - **`types/`**: Data types/models.
    - **`middleware/`**: Middleware for various functionalities.

4. **`test/`**: Integration tests, Load tests.

4. **`docs/`**: Swagger generated UI.

## Getting Started

### Prerequisites

- Go 1.21+
- Docker
- Docker Compose
- Make (Optional, but highly recommended)

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

To run the project localy, use `make setup` to install swag and all other necessary dependencies. Use `go run main.go` to run the application. This will run the decimal-to-roman-numerals. To run all the containers, i.e. the application, prometheus and grafana use the command `make up`.

Here are a list of commands to run the application via containers.

| Command       | Description                                                                 |
|---------------|-----------------------------------------------------------------------------|
| `make setup`  | Installs Swag for API documentation, initializes Swag documentation, and builds the Go project. |
| `make build`  | Builds Docker images without using cache.                                   |
| `make test`   | Runs the tests using Docker Compose.                                        |
| `make cover`  | Runs tests and generates a coverage report inside a Docker container.       |
| **`make up`**     | **Starts Docker containers for the application, Prometheus, and Grafana.**      |
| `make down`   | Stops and removes Docker containers.                                        |
| `make restart`| Restarts Docker containers.                                                 |
| `make clean`  | Stops and removes Docker containers and images, prunes Docker volumes, and removes build artifacts. |

## API Documentation

The API is documented using Swagger and can be accessed at `http://localhost:8001/swagger/index.html`.

### Endpoints

#### 1. Convert Number(s) to Roman Numerals

This endpoint converts a comma-separated list of numbers to their corresponding Roman numeral representations.

- **URL**: `/api/v1/convert`
- **Method**: `GET`
- **Parameters**:
  - `numbers` (required): Comma-separated list of integers to be converted. Each number must be within the range 1 to 3999.
- **Example**: `/api/v1/convert?numbers=10,50,100`

#### Response
- **Status Code**: `200 OK`
- **Body**: JSON object containing the results.
  - `results`: An array of objects containing the decimal number and its Roman numeral representation.
    - `number`: Decimal number.
    - `roman`: Roman numeral representation.

#### Example
Request:
```http
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

#### 2. Convert Range(s) of Numbers to Roman Numerals

This endpoint converts multiple ranges of numbers to their corresponding Roman numeral representations.

- **URL**: `/api/v1/convert`
- **Method**: `POST`
- **Body**: JSON object containing an array of number ranges.
  - `ranges`: An array of objects specifying number ranges.
    - `min`: The minimum value of the range *(inclusive)*.
    - `max`: The maximum value of the range (inclusive)*.
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
    {"min": 10, "max": 12},
    {"min": 50, "max": 53}
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
    {"number": 50, "roman": "L"},
    {"number": 51, "roman": "LI"},
    {"number": 52, "roman": "LII"},
    {"number": 53, "roman": "LIII"}
  ]
}
```

## Logging And Monitoring

Docker compose handles the integration of prometheus and grafana instances using the provided config files.

### 1. Prometheus

To expose various metrics in a Go application, we have provided a `/metrics` HTTP endpoint. We can view all the exposed metrics via this url `http://localhost:8001/metrics`. This consumed by `Prometheus` and finally visualized by `grafana`.

### 2. Grafana

This project has grafana integration for visualisation of the metrics.

| URL                                  | Description                                       |
|--------------------------------------|---------------------------------------------------|
| [http://localhost:9090](http://localhost:9090) | Link to the Prometheus instance                   |
| [http://localhost:3000](http://localhost:3000) | Link to the Grafana instance (login: `admin`, password: `admin`) |


There are two dashboards included. The `Gin Application Metrics` dashboard captures various statistics like total requests, request rates, status code distribution etc.

| URL                                  | Description                                       |
|--------------------------------------|---------------------------------------------------|
| [Gin Application Metrics](http://localhost:3000/d/FDB061FMz/gin-application-metrics?orgId=1&refresh=5s) | Dashboard showing stats like total requests, status code distribuion etc.                   |
| [Go Metrics](http://localhost:3000/d/CgCw8jKZz/go-metrics?orgId=1&refresh=5s) | Dashboard showing various go memory stats |


[![apistats.png](https://i.postimg.cc/vHnR15Nv/Screenshot-2024-05-18-230423.png)](https://postimg.cc/HVTvR8HJ)

The second dashboard covers various go statistics.

[![gostats.png](https://i.postimg.cc/wMSSfB7J/gostats.png)](https://postimg.cc/9wY8zCPF)


## Testing
| Command         | Description                                     |
|-----------------|-------------------------------------------------|
| `go test ./...` | Runs all the tests in the local environment     |
| `make test`     | Runs all the tests using Docker                 |
| `make cover`    | Generates the test coverage report using Docker |

NOTE: The Make file has been tested in the Windows environment. Although The makefile is cross-platform compatible, there might some small errors in Unix based systems.

### 1. Unit tests
This project has extensive unit tests covering. The test files sit next to the code files with a suffix of `_test.go`.

### 2. Integration tests

The integration tests for this project are designed to ensure that the API endpoints are functioning correctly. These tests cover various scenarios including valid inputs, invalid inputs, edge cases, and performance under load.

#### Test Setup

The tests are written using the Go testing package and the Gin web framework. The `SetupRouter` function initializes the Gin router for testing purposes.

#### Helper Functions

- **checkStatus**: Verifies that the response status code matches the expected status.
- **checkResponse**: Unmarshals the response body and checks the result for GET requests.
- **checkPostResponse**: Unmarshals the response body and checks the result for POST requests.

#### Test Cases

##### GET /api/v1/convert

1. **Valid Inputs**: Tests the conversion of various integers to Roman numerals.
    - Example: `1` to `I`, `58` to `LVIII`, `3999` to `MMMCMXCIX`.

2. **Invalid Inputs**: Tests the response for invalid inputs such as non-numeric strings and out-of-range values.
    - Example: `"abc"`, `-1`, `4000`.

3. **Edge Cases**: Tests the boundary values for valid Roman numeral conversions.
    - Example: `1` to `I`, `3999` to `MMMCMXCIX`.

4. **Performance**: Tests the performance of the handler under a load of 1000 requests.

##### POST /convert

1. **Valid Inputs**: Tests the conversion of ranges of integers to Roman numerals.
    - Example payload: `{"ranges": [{"min": 10, "max": 12}, {"min": 14, "max": 16}]}`.
    - Expected results: `10` to `X`, `11` to `XI`, `12` to `XII`, `14` to `XIV`, `15` to `XV`, `16` to `XVI`.

2. **Invalid Inputs**: Tests the response for various invalid range inputs.
    - Example: Empty ranges, non-integer values, negative values, max less than min.

3. **Edge Cases**: Tests various edge cases for the range inputs.
    - Example: Single number range, very small range, maximum valid range, overlapping large ranges, reverse order ranges.


### 3. Load Testing

The load tests for this project are designed to evaluate the performance and reliability of the API endpoints under high request volumes. These tests simulate concurrent requests to ensure the system can handle heavy traffic.

### Load Test Setup

Each test distributes a total of 1000 requests across multiple goroutines to simulate concurrent users.

### Helper Functions

- **performRequest**: Executes a GET request and returns the response.
- **performPostRequest**: Executes a POST request with a given payload and returns the response.

### Test Cases

#### GET /api/v1/convert Load Test

- **TestConvertHandlerLoad**: This test performs 1000 GET requests to the `/api/v1/convert` endpoint with the number `123`. It checks the response status code and validates the response body to ensure the correct conversion to Roman numerals.

  **Validation**:
  - Status Code: Should be `200 OK`.
  - Response Body: Should contain the correct conversion result, e.g., `123` to `CXXIII`.

#### POST /convert Load Test

- **TestConvertRangesHandlerLoad**: This test performs 1000 POST requests to the `/convert` endpoint with a payload containing ranges of numbers to convert to Roman numerals. It checks the response status code and validates the response body for correctness.

  **Payload**:
  ```json
  {
      "ranges": [
          {"min": 10, "max": 15},
          {"min": 20, "max": 25}
      ]
  }
  ```

**Expected Results**:
- Numbers from 10 to 15: X, XI, XII, XIII, XIV, XV
- Numbers from 20 to 25: XX, XXI, XXII, XXIII, XXIV, XXV

**Validation:**
- Status Code: Should be 200 OK.
- Response Body: Should correctly reflect the expected Roman numeral conversions for the provided ranges.

## TODOs

- [ ] Authentication!
- [ ] Rate Limiter!
- [ ] Load balancer!
- [ ] AutoScaling!
- [x] Redis DB support: Caching is not necessary for this project. Won't do.
