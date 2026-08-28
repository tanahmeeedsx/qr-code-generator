# QR Code Generator API

A simple REST API built with Go that generates QR codes from text or URLs and returns them as PNG images.

## Features

* Generate QR codes from text or URLs
* Simple REST API
* Health check endpoint
* Custom QR code size
* Unit tests
* GitHub Actions CI
* No Docker required

## Technologies

* Go
* `net/http`
* `github.com/skip2/go-qrcode`
* GitHub Actions

## Project Structure

```text
qr-code-generator/
├── .github/
│   └── workflows/
│       └── ci.yml
├── go.mod
├── go.sum
├── main.go
├── main_test.go
└── README.md
```

## Prerequisites

* Go 1.25 or later

Check your installed Go version:

```bash
go version
```

## Run Locally

Clone the repository:

```bash
git clone git@github.com:tanahmeeedsx/qr-code-generator.git
cd qr-code-generator
```

Download dependencies:

```bash
go mod download
```

Start the server:

```bash
go run main.go
```

The API will run at:

```text
http://localhost:8080
```

## API Endpoints

### GET `/health`

Checks whether the API is running.

```bash
curl http://localhost:8080/health
```

Example response:

```text
QR Code Generator is running!
```

### POST `/generate`

Generates a QR code from text or a URL.

#### Generate from Text

```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"text":"Hello, World!"}' \
  --output qrcode.png
```

#### Generate from URL

```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}' \
  --output qrcode.png
```

#### Generate with Custom Size

```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com","size":300}' \
  --output qrcode.png
```

The generated QR code is returned as a PNG image.

## Testing

Run the test suite:

```bash
go test -v ./...
```

## Build

Build the application:

```bash
go build -v .
```

## CI

This project uses GitHub Actions for continuous integration.

The CI workflow runs when:

* Code is pushed to the `main` branch
* A pull request targets the `main` branch

The workflow:

1. Checks out the repository
2. Sets up Go
3. Downloads dependencies
4. Runs the test suite
5. Builds the application

## Docker

Docker is **not used** in this project.

The application runs directly as a Go application on the host or server.
