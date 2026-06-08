package httpapi

import "net/http"

type statusResponse struct {
	Status string `json:"status"`
}

func (app *Application) statusCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.methodNotAllowed(w, http.MethodGet)
		return
	}

	app.writeJSON(w, http.StatusOK, statusResponse{Status: "ok"})
}
