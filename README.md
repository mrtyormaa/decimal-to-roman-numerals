# decimal-to-roman-numerals

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
    - **`database/`**: Database connection and queries.

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

### API Documentation

The API is documented using Swagger and can be accessed at:

```
http://localhost:8001/swagger/index.html
```

## Usage

### Endpoints

- `GET /`: Health Check.
- `GET /GetRoman`: Dummy Test Endpoint.

## TODOs

- Endpoint Implementations
- Possibly Authentication!
- Possible Rate Limiter!
- Possibly Redis DB support!
