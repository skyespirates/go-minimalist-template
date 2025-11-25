package main

import "net/http"

func (app *application) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.LogInfo(r, "request")
		next.ServeHTTP(w, r)
		app.logger.LogInfo(r, "response")
	})
}
