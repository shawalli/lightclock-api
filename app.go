// app.go

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

type Response struct {
	Error	string		`json:"error"`
	Result	interface{}	`json:"result"`
}

func (a *App) Init() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run() { }

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSONFull(w, code, nil, message)
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	respondJSONFull(w, code, payload, "")
}

func respondJSONFull(w http.ResponseWriter, code int, payload interface{}, message string) {
	resp := Response{
		Error:	message,
		Result:	payload,
	}
	
	response, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/event", a.createEvent).Methods("POST")
}