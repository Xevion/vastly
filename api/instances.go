package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InstanceGetResponse represents the response from GET /instances
type InstanceGetResponse struct {
	Instances []Instance `json:"instances"`
}

// GetInstances retrieves all instances for the authenticated user
func (c *Client) GetInstances() (*InstanceGetResponse, error) {
	resp, err := c.makeRequest(http.MethodGet, "/instances", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result InstanceGetResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteInstance removes an instance by ID
func (c *Client) DeleteInstance(id int) error {
	path := fmt.Sprintf("/instances/%d", id)
	resp, err := c.makeRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}

// PutInstance updates an instance's status
func (c *Client) PutInstance(id int, status string) error {
	path := fmt.Sprintf("/instances/%d", id)
	data := map[string]string{"status": status}

	resp, err := c.makeRequest(http.MethodPut, path, data)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
