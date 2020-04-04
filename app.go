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

func (a *App) Init() {
	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run() { }

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/event", a.createEvent).Methods("POST")
}