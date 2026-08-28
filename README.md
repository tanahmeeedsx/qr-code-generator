# QR Code Generator API

A simple REST API built with Go that generates QR codes from text or URLs and returns them as PNG images.

## Features

* Generate QR codes from text or URLs
* Health check endpoint
* Custom QR code size
* Unit tests
* GitHub Actions CI

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

```bash
go version
```

## Run Locally

```bash
git clone git@github.com:tanahmeeedsx/qr-code-generator.git
cd qr-code-generator
go mod download
go run main.go
```

API: `http://localhost:8080`

## API

### GET `/health`

```bash
curl http://localhost:8080/health
```

Response:

```text
QR Code Generator is running!
```

### POST `/generate`

Generate a QR code from text:

```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"text":"Hello, World!"}' \
  --output qrcode.png
```

Generate a QR code from a URL:

```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}' \
  --output qrcode.png
```

Custom size:

```bash
curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com","size":300}' \
  --output qrcode.png
```

## Testing

```bash
go test -v ./...
```

## Build

```bash
go build -v .
```

## CI

GitHub Actions automatically runs tests and builds the application on pushes to `main` and pull requests targeting `main`.
