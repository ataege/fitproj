# FitProj

A Go application built with the Fiber web framework.

## Project Structure

```
├── api/        # API documentation, OpenAPI specs
├── cmd/        # Application entry points
│   └── main.go # Main application entry point
├── config/     # Configuration related code
├── docs/       # Documentation
├── internal/   # Private application code
│   ├── handlers/   # HTTP handlers
│   ├── middleware/ # Middleware components
│   └── routes/     # Route definitions
├── pkg/        # Public libraries that can be used by external applications
```

## Getting Started

### Prerequisites

- Go 1.19 or higher

### Installation

1. Clone the repository
2. Install dependencies:

```bash
go mod tidy
```

### Running the Application

```bash
go run cmd/main.go
```

Or build and run:

```bash
go build -o fitproj cmd/main.go
./fitproj
```

## Environment Variables

Copy the `.env.example` to `.env` and adjust the values as needed:

```bash
cp .env.example .env
```

## API Documentation

API endpoints and documentation can be found in the `api/` directory.

## License

This project is licensed under the MIT License.
