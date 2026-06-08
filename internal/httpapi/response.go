package httpapi

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func (app *Application) writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		app.logger.Printf("write json response: %v", err)
	}
}

func (app *Application) writeError(w http.ResponseWriter, status int, message string) {
	app.writeJSON(w, status, errorResponse{Error: message})
}

func (app *Application) methodNotAllowed(w http.ResponseWriter, allowedMethods ...string) {
	for _, method := range allowedMethods {
		w.Header().Add("Allow", method)
	}
	app.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
}
