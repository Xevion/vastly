package api

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	client := NewClient(apiKey)

	if client.apiKey != apiKey {
		t.Errorf("expected apiKey to be %s, got %s", apiKey, client.apiKey)
	}

	if client.baseURL != "https://console.vast.ai/api/v0" {
		t.Errorf("expected baseURL to be %s, got %s", "https://console.vast.ai/api/v0", client.baseURL)
	}

	if client.httpClient == nil {
		t.Error("expected httpClient to be initialized")
	}
}

func TestApply(t *testing.T) {
	apiKey := "abc123"
	client := NewClient(apiKey)
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	client.Apply(req)

	expected := "Bearer " + apiKey
	if got := req.Header.Get("Authorization"); got != expected {
		t.Errorf("expected Authorization header to be %s, got %s", expected, got)
	}
}
