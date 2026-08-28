package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/skip2/go-qrcode"
)

// GenerateRequest defines the JSON payload schema for the /generate endpoint.
type GenerateRequest struct {
	Text string `json:"text"`
	URL  string `json:"url"`
	Size int    `json:"size,omitempty"`
}

// healthHandler handles GET /health requests.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "QR Code Generator is running!")
}

// generateHandler handles POST /generate requests and returns a PNG QR code image.
func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Support both 'text' and 'url' fields
	content := req.Text
	if content == "" {
		content = req.URL
	}
	content = strings.TrimSpace(content)

	if content == "" {
		http.Error(w, "'text' or 'url' is required", http.StatusBadRequest)
		return
	}

	size := req.Size
	if size <= 0 {
		size = 256
	}

	// Generate QR code as PNG image bytes
	pngBytes, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate QR code: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(pngBytes)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/generate", generateHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

