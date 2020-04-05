// handlers_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

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
