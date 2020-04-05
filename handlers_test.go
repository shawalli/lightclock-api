// handlers_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventPost(t *testing.T) {
	var requestBody = []byte(`{"label":"morning"}`)
	
	req, _ := http.NewRequest("POST", "/event", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	assert.Equal(t, http.StatusCreated, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Empty(t, respBody["error"])

	assert.NotNil(t, respBody["result"])

	result := respBody["result"].(map[string]interface{})

	assert.Equal(t, "morning", result["label"])

	assert.NotNil(t, result["id"])
}

func TestEventPostInvalidJSON(t *testing.T) {
	var requestBody = []byte(`{"label"}`)

	req, _ := http.NewRequest("POST", "/event", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t, "Invalid request payload", respBody["error"])

	assert.Nil(t, respBody["result"])
}

func TestEventPostMissingFields(t *testing.T) {
	var requestBody = []byte(`{}`)
	
	req, _ := http.NewRequest("POST", "/event", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t, "Missing required fields", respBody["error"])

	assert.Nil(t, respBody["result"])
}
