// handlers.go

package main

import (
	"encoding/json"
	"net/http"
)

func (a *App) createEvent(w http.ResponseWriter, r *http.Request) {
	var evt event
	decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&evt); err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
        return
	}
	defer r.Body.Close()

	if err := evt.createEvent(); err != nil {
		respondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, http.StatusCreated, evt)
}

// func registerEventHandlers(a* App) {

// }