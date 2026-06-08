package httpapi

import "net/http"

func (app *Application) Routes() http.Handler {
	var handler http.Handler = http.HandlerFunc(app.routeRequest)
	handler = app.requestTimeout(handler)
	handler = app.securityHeaders(handler)
	handler = app.logRequest(handler)
	handler = app.recoverPanic(handler)

	return handler
}

func (app *Application) routeRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/v1/status":
		app.statusCheckHandler(w, r)
	default:
		app.writeError(w, http.StatusNotFound, "not found")
	}
}
