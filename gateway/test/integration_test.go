package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var gatewayURL string = "http://127.0.0.1:8080"

func TestCreateHashEndpoint(t *testing.T) {
	payload := map[string]string{"payload": "your_payload_here"}
	payloadBytes, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post(gatewayURL+"/createhash", "application/json", bytes.NewBuffer(payloadBytes))
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	hash := response["hash"].(string)
	assert.NotEmpty(t, hash)
}

func TestCheckHashEndpoint(t *testing.T) {
	payload := map[string]string{"payload": "your_payload_here"}
	payloadBytes, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post(gatewayURL+"/checkhash", "application/json", bytes.NewBuffer(payloadBytes))
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	exists := response["exists"].(bool)
	assert.True(t, exists)
}

func TestGetHashEndpoint(t *testing.T) {
	payload := map[string]string{"payload": "your_payload_here"}
	payloadBytes, err := json.Marshal(payload)
	assert.NoError(t, err)

	resp, err := http.Post(gatewayURL+"/gethash", "application/json", bytes.NewBuffer(payloadBytes))
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	hash := response["hash"].(string)
	assert.NotEmpty(t, hash)
}
