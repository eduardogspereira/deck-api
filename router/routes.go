package router

import (
	"net/http"

	"github.com/eduardogspereira/deck-api/router/healthcheck"
	"github.com/gorilla/mux"
)

// NewHTTPHandler create a new router to handle HTTP requests to the server.
func NewHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/_health", healthcheck.Handler)

	return router
}
