package handlers

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	response, err := http.Get("http://localhost:8080/api/v1/health")

	if err != nil {
		t.Fatalf("failed to call health endpoint: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			response.StatusCode,
		)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("failed to read response: %v", err)
	}

	if !strings.Contains(string(body), "healthy") {
		t.Fatalf(
			"expected response to contain 'healthy', got %s",
			string(body),
		)
	}
}