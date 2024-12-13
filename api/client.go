package api

import "net/http"

// Client handles API communication with the Vast.ai API
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// Error represents an API error response
type APIError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Msg     string `json:"msg"`
}

// NewClient creates a new Vast.ai API client
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		baseURL:    "https://console.vast.ai/api/v0",
		httpClient: &http.Client{},
	}
}

// Apply applies the Bearer token to the request
func (c *Client) Apply(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
}

// makeRequest creates and sends an HTTP request with the provided method, path and body
func (c *Client) makeRequest(method, path string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	c.Apply(req)
	return c.httpClient.Do(req)
}
