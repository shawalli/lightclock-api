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

func TestEventGet(t *testing.T) {
	var requestBody = []byte(`{"label":"morning"}`)
	
	req, _ := http.NewRequest("POST", "/event", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp := executeRequest(req)
	assert.Equal(t, http.StatusCreated, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t, "morning", respBody["label"])

	assert.NotNil(t, respBody["id"])
}