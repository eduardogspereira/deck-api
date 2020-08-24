package healthcheck

import (
	"net/http"
)

// Handler for the health check route
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
