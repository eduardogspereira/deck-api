package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/eduardogspereira/deck-api/router/healthcheck"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/_health", healthcheck.Handler)
	return r
}

var _ = Describe("HealthCheck", func() {
	var router = setupRouter()

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_health", nil)
	router.ServeHTTP(rr, req)

	It("should have returned status_code 200", func() {
		Expect(rr.Code).To(Equal(http.StatusOK))
	})

	It("should have returned OK in body", func() {
		Expect(rr.Body.String()).To(Equal("OK"))
	})
})

func TestHealthCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HealthCheck Suite")
}
