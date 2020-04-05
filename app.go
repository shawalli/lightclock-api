// app.go

package main

import (
	"encoding/json"
	"net/http"
	"fmt"

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

	a.Router.Use(loggingMiddleware)

	a.initializeRoutes()
}

func (a *App) Run() {
	port := 8080

	address := fmt.Sprintf(":%d", port)

	logInfo("Serving on %v\n", address)

	http.ListenAndServe(address, a.Router)
}

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
	a.Router.HandleFunc("/", a.GetIndex).Methods("GET")
	a.Router.HandleFunc("/event", a.createEvent).Methods("POST")
}