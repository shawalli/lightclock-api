// handlers.go

package main

import (
	"encoding/json"
	"net/http"
)

func (a *App) GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte("This is the lightclock API"))
}

func (a *App) createEvent(w http.ResponseWriter, r *http.Request) {
	var evt event
	decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&evt); err != nil {
        respondError(w, http.StatusBadRequest, "Invalid request payload")
        return
	}
	defer r.Body.Close()

	if evt.Label == "" {
        respondError(w, http.StatusBadRequest, "Missing required fields")
        return
	}

	if err := evt.createEvent(); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, evt)
}
