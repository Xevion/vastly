//go:build real

package api

import (
	"testing"
)

func GetInstancesTest(t *testing.T) {
	// Create a new client with the test server's URL
	client := NewClient("test")

	// Call the GetInstances method
	resp, err := client.GetInstances()
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the response is valid
	if len(resp.Instances) != 0 {
		t.Fatalf("expected 0 instances, got %d", len(resp.Instances))
	}
}
