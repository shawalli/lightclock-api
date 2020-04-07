// app_test.go

package main

import (
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

func TestInvalidPath(t *testing.T) {
	req, _ := http.NewRequest("GET", "/foobar", nil)

	resp := executeRequest(req)
	assert.Equal(t, http.StatusNotFound, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t, respBody["error"], "Not Found")

	assert.Nil(t, respBody["result"])
}

func TestInvalidMethod(t *testing.T) {
	a.Router.HandleFunc("/", a.GetHealthCheck).Methods("GET")

	req, _ := http.NewRequest("POST", "/", nil)

	resp := executeRequest(req)
	assert.Equal(t, http.StatusNotFound, resp.Code)

	var respBody map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t, respBody["error"], "Not Found")

	assert.Nil(t, respBody["result"])
}
