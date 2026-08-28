package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", res.StatusCode)
	}

	body := w.Body.String()
	expected := "QR Code Generator is running!\n"
	if body != expected {
		t.Fatalf("expected body %q, got %q", expected, body)
	}
}

func TestHealthHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	w := httptest.NewRecorder()

	healthHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405 Method Not Allowed, got %d", res.StatusCode)
	}
}

func TestGenerateHandler_TextSuccess(t *testing.T) {
	payload := []byte(`{"text": "https://example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	generateHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "image/png" {
		t.Fatalf("expected Content-Type image/png, got %s", contentType)
	}

	body := w.Body.Bytes()
	// Check PNG magic bytes: \x89PNG\r\n\x1a\n
	pngHeader := []byte{0x89, 'P', 'N', 'G', 0x0D, 0x0A, 0x1A, 0x0A}
	if len(body) < 8 || !bytes.Equal(body[:8], pngHeader) {
		t.Fatalf("response body is not a valid PNG image")
	}
}

func TestGenerateHandler_URLSuccess(t *testing.T) {
	payload := []byte(`{"url": "https://golang.org", "size": 128}`)
	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	generateHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "image/png" {
		t.Fatalf("expected Content-Type image/png, got %s", contentType)
	}
}

func TestGenerateHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/generate", nil)
	w := httptest.NewRecorder()

	generateHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected status 405 Method Not Allowed, got %d", res.StatusCode)
	}
}

func TestGenerateHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/generate", strings.NewReader("invalid json"))
	w := httptest.NewRecorder()

	generateHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400 Bad Request, got %d", res.StatusCode)
	}
}

func TestGenerateHandler_MissingContent(t *testing.T) {
	payload := []byte(`{"text": "   "}`)
	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	generateHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status 400 Bad Request, got %d", res.StatusCode)
	}
}
