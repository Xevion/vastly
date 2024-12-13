package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// Client handles API communication with the Vast.ai API
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	logger     *zap.SugaredLogger
}

// Error represents an API error response
type APIError struct {
	Success bool `json:"success"`
}

// NewClient creates a new Vast.ai API client
func NewClient(apiKey string) *Client {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	return &Client{
		apiKey:     apiKey,
		baseURL:    "https://console.vast.ai/api/v0",
		httpClient: &http.Client{},
		logger:     sugar,
	}
}

// Apply applies the Bearer token to the request
func (c *Client) Apply(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
}

// makeRequest creates and sends an HTTP request with the provided method, path and body
func (c *Client) makeRequest(method, path string, body interface{}) (*http.Response, error) {
	c.logger.Debugw("making request",
		"method", method,
		"path", path,
		"body", body,
	)

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		c.logger.Errorw("failed to create request",
			"error", err,
			"method", method,
			"path", path,
		)
		return nil, err
	}

	c.Apply(req)
	res, err := c.httpClient.Do(req)

	if err != nil {
		c.logger.Errorw("failed to send request",
			"error", err,
			"method", method,
			"path", path,
		)
		return nil, err
	}

	// debug print response
	res_body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var apiError APIError
		if err := json.NewDecoder(bytes.NewBuffer(res_body)).Decode(&apiError); err != nil {
			c.logger.Errorw("failed to decode error response",
				"error", err,
			)
		} else {
			c.logger.Errorw("api error",
				"error", apiError,
			)
		}
	} else {
		c.logger.Debugw("response",
			"status", res.Status,
			"status_code", res.StatusCode,
		)
	}

	res.Body = io.NopCloser(bytes.NewBuffer(res_body))

	return res, nil
}
