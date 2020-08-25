package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardogspereira/deck-api/router/healthcheck"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/_health", healthcheck.Handler)
	return r
}

func TestHealthCheckHandler(t *testing.T) {
	router := setupRouter()

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_health", nil)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedBody := "OK"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}
