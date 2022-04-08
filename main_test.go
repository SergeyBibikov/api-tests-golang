package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTheServiceIsReady(t *testing.T) {
	client := GetClient()

	resp, err := client.R().Get("/ready")
	if err != nil {
		t.Errorf("The service is not ready %s", err)
	}

	var respB Readiness
	json.Unmarshal(resp.Body(), &respB)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, respB.Status, "ready")
}
