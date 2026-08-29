# QR Code Generator API

A simple REST API built with Go that generates QR codes from text or URLs and returns them as PNG images.

## Features

- Generate QR codes from text or URLs
- Health check endpoint
- Custom QR code size
- Unit tests
- GitHub Actions CI

## Technologies

- Go
- `net/http`
- `github.com/skip2/go-qrcode`
- GitHub Actions

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
