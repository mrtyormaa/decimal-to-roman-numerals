# decimal-to-roman-numerals

[![Test and coverage](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions/workflows/ci.yml/badge.svg)](https://github.com/mrtyormaa/decimal-to-roman-numerals/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals/graph/badge.svg?token=WCPsoNnQEy)](https://codecov.io/github/mrtyormaa/decimal-to-roman-numerals)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtyormaa/decimal-to-roman-numerals)](https://goreportcard.com/report/github.com/mrtyormaa/decimal-to-roman-numerals)
[![Maintainability](https://api.codeclimate.com/v1/badges/dfbf91b073b8fec1f6bf/maintainability)](https://codeclimate.com/github/mrtyormaa/decimal-to-roman-numerals/maintainability)

## Overview

This repository provides an API endpoint to convert decimal number ranges to Roman numerals. The next steps/thoughts to take the project to production can be found at [production.md](./production.md).

**TL;DR to Run and See Results**

Execute `make up` command and visit the links below.

| Link                             | Description                                                |
|----------------------------------|------------------------------------------------------------|
| [http://localhost:8001/swagger/index.html](http://localhost:8001/swagger/index.html) | Link to the Swagger API documentation                  |
| [http://localhost:8001/metrics](http://localhost:8001/metrics) | Link to the metrics endpoint used for Prometheus                |
| [http://localhost:9090](http://localhost:9090) | Link to the Prometheus instance                        |
| [http://localhost:3000](http://localhost:3000) | Link to the Grafana instance, login: `admin`, password: `admin` |
| [Gin Application Metrics](http://localhost:3000/d/FDB061FMz/gin-application-metrics?orgId=1&refresh=5s) | Grafana Dashboard showing stats like total requests, status code distribuion etc.                   |
| [Go Metrics](http://localhost:3000/d/CgCw8jKZz/go-metrics?orgId=1&refresh=5s) | Grafana Dashboard showing various go memory stats |


## Folder structure
```
decimal-to-roman-numerals/
|-- bin/                        # Contains the compiled binaries
|-- config/                     # Configuration files and settings
|-- docs/                       # Documentation and Swagger UI
|   |-- docs.go                 # Go file for documentation generation
|   |-- swagger.json            # Swagger specification in JSON format
|   |-- swagger.yaml            # Swagger specification in YAML format
|-- pkg/                        # Libraries and packages for external use
|   |-- api/                    # API logic
|       |-- roman/              # Roman numeral conversion logic
|           |-- handler.go      # HTTP handlers for Roman numeral conversion
|       |-- router.go           # API routes
|   |-- types/                  # Data types and models
|   |-- middleware/             # Middleware for various functionalities
|-- test/                       # Integration and load tests
|-- Dockerfile                  # Dockerfile for building the container
|-- docker-compose.yml          # Docker Compose file for multi-container applications
|-- go.mod                      # Go module file
|-- go.sum                      # Go dependencies checksum file
|-- main.go                     # The entry point of the application
|-- README.md                   # Project overview and instructions
|-- Makefile                    # Makefile for building, testing, and running the application
``` 
[![uml-class-diagram.png](https://i.postimg.cc/76CzGSjD/uml-class-diagram.png)](https://postimg.cc/bDhr4SzF)

## Getting Started

### Prerequisites

- Go 1.21+
- Docker
- Docker Compose
- Make (Optional, but highly recommended)

### Installation and Running

1. Clone the repository

    ```bash
    git clone https://github.com/mrtyormaa/decimal-to-roman-numerals
    ```

2. Navigate to the directory

    ```bash
    cd decimal-to-roman-numerals
    ```

3. Set Enviroment Variables (Optional)

Create an `.env` file with following configurations. If this step is not done, the default values will be used as shown below:

    ```
    PORT=8001
    GF_SECURITY_ADMIN_USER=admin
    GF_SECURITY_ADMIN_PASSWORD=admin
    ```

3. Build and run the Docker containers

This setup defines a multi-stage build process in the Dockerfile, ensuring separation of concerns during the `build`, `testing`, `coverage`, and final application stages. The `docker-compose.yml` file integrates `Prometheus` and `Grafana` as optional dependencies, allowing you to monitor your application effectively.

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

### Make Optional commands
**Docker Compose Commands for `decimal-to-roman-numerals`**

| **Operation**              | **Command**                                                                                                              |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------|
| **Setup**                  |                                                                                                                          |
| Install Swag               | `go install github.com/swaggo/swag/cmd/swag@latest`                                                                      |
| Download Go Modules        | `go mod download`                                                                                                        |
| Initialize Swag Documentation | `swag init`                                                                                                             |
| Build Go Project           | `go build -o bin/main main.go`                                                                                           |
| **Build Docker Images**    |                                                                                                                          |
| Build without cache        | `docker-compose -f docker-compose.yml build --no-cache`                                                                  |
| **Test the Application**   |                                                                                                                          |
| Run tests                  | `docker-compose -f docker-compose.yml up --build roman-numerals-tests`                                                   |
| **Generate Coverage Report** |                                                                                                                        |
| Build coverage container   | `docker build -t decimal-to-roman-numerals-coverage -f Dockerfile --target coverage .`                                   |
| Run coverage container     | `docker run --rm -v $(pwd)/coverage:/coverage decimal-to-roman-numerals-coverage`                                        |
| **Start Services**         |                                                                                                                          |
| Start application, Prometheus, and Grafana | `docker-compose -f docker-compose.yml up -d roman-numerals prometheus grafana`                                          |
| **Stop Services**          |                                                                                                                          |
| Stop and remove containers | `docker-compose -f docker-compose.yml down`                                                                              |
| **Restart Services**       |                                                                                                                          |
| Restart containers         | `docker-compose -f docker-compose.yml restart`                                                                           |
| **Clean Up**               |                                                                                                                          |
| Stop and remove containers, images, and volumes | `docker-compose -f docker-compose.yml down --rmi all`<br>`docker volume prune -f`                           |


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
GET /api/v1/convert?numbers=10,50,100
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
POST /api/v1/convert
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

This project has grafana integration for visualisation of the metrics. If you want to change the default id(admin) and password(admin), create an `.env` file and give values to the following variables:
```
GF_SECURITY_ADMIN_USER=
GF_SECURITY_ADMIN_PASSWORD=
```

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

1. **Valid Inputs**: Tests the conversion of various integers to Roman numerals. The api handles leading zeroes, leading `+` sign and extra spaces.
    - Example: `1` to `I`, `58` to `LVIII`, `3999` to `MMMCMXCIX`.

2. **Invalid Inputs**: Tests the response for invalid inputs such as non-numeric strings, out-of-range values, all non-numeric unicodes, arithmatic operations.
    - Example: `"abc"`, `-1`, `4000`.

3. **Edge Cases**: Tests the boundary values for valid Roman numeral conversions.
    - Example: `1` to `I`, `3999` to `MMMCMXCIX`.

4. **Performance**: Tests the performance of the handler under a load of 1000 requests.

4. **Correctness**: Tests the conversion of various integers to Roman numerals using a different algorithm.

##### POST /convert

1. **Valid Inputs**: Tests the conversion of ranges of integers to Roman numerals.
    - Example payload: `{"ranges": [{"min": 10, "max": 12}, {"min": 14, "max": 16}]}`.
    - Expected results: `10` to `X`, `11` to `XI`, `12` to `XII`, `14` to `XIV`, `15` to `XV`, `16` to `XVI`.

2. **Invalid Inputs**: Tests the response for various invalid range inputs.
    - Example: Empty ranges, non-integer values, negative values, max less than min, invalid jsons etc. Leading zeroes, leading `+` sign are not supported here.

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

## CI/CD Process

### Automated Release

This workflow is designed to automate the creation of GitHub Releases. It adheres to [Semantic Versioning](https://semver.org/), which is a versioning scheme that uses a three-part version number: `MAJOR.MINOR.PATCH`. This workflow is triggered whenever a commit tag that starts with "v" (e.g., "v1.0.0", "v0.1.4") is pushed to the repository. The workflow performs the following steps:

- **Trigger:** Executes on every push to a tag that matches the pattern "v*".
- **Job Name:** Release
- **Environment:** Runs on the latest Ubuntu environment.
- **Action:** Utilizes the `marvinpinto/action-automatic-releases` action to create a GitHub Release with the details of the relevant commits.
- **Authentication:** Uses the GitHub token stored in the repository secrets to authenticate the release creation.

### Test and Coverage

This project uses GitHub actions to automate the tests upon push or pull request to main branch. This workflow ensures code quality by running tests and collecting coverage data. It is triggered on every push and pull request. The workflow performs the following steps:

- **Trigger:** Executes on every push to the repository and on every pull request.
- **Job Name:** Build
- **Environment:** Runs on the latest Ubuntu environment.
- **Steps:**
  - **Checkout Code:** Uses the `actions/checkout` action to check out the repository code.
  - **Set Up Go Environment:** Uses the `actions/setup-go` action to set up a Go environment with the latest stable version.
  - **Gather Dependencies:** Downloads the Go module dependencies.
  - **Run Tests and Coverage:** Executes tests with race condition detection and generates a coverage profile.
  - **Upload Coverage:** Uses the `codecov/codecov-action` to upload the coverage report to Codecov, using the Codecov token stored in the repository secrets.

### Deployment - In Progress
This is currently in progress under `feature/kubernetesIntegration` branch.


