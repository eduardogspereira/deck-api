package deck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardogspereira/deck-api/router/deck"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func routerFactory() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	return router
}

var _ = Describe("Deck Hanlders", func() {
	Describe("Create", func() {
		Context("When creating a Deck with default options", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{id: "ABC-DEF"}
			router.POST("/deck", deck.CreateBuilder(deckRepo))

			req, _ := http.NewRequest("POST", "/deck", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 201", func() {
				Expect(rr.Code).To(Equal(201))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"deck_id":"ABC-DEF","shuffled":false,"remaining":52}`),
				)
			})
		})

		Context("When deckRepo.Save returns an error", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{returnErrorOnSave: true}
			router.POST("/deck", deck.CreateBuilder(deckRepo))

			req, _ := http.NewRequest("POST", "/deck", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 500", func() {
				Expect(rr.Code).To(Equal(500))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"error":"internal error"}`),
				)
			})
		})

		Context("When deck created with shuffle option", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{id: "ABC-DEF"}
			router.POST("/deck", deck.CreateBuilder(deckRepo))

			req, _ := http.NewRequest("POST", "/deck?shuffle=true", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 201", func() {
				Expect(rr.Code).To(Equal(201))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"deck_id":"ABC-DEF","shuffled":true,"remaining":52}`),
				)
			})
		})

		Context("When deck requested with invalid shuffle option", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{id: "ABC-DEF"}
			router.POST("/deck", deck.CreateBuilder(deckRepo))

			req, _ := http.NewRequest("POST", "/deck?shuffle=indianajones", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 400", func() {
				Expect(rr.Code).To(Equal(400))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"error":"invalid request"}`),
				)
			})
		})

		Context("When deck created with cards option", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{id: "ABC-DEF"}
			router.POST("/deck", deck.CreateBuilder(deckRepo))

			req, _ := http.NewRequest("POST", "/deck?cards=AS,KS,AC,2C,KH", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 201", func() {
				Expect(rr.Code).To(Equal(201))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"deck_id":"ABC-DEF","shuffled":false,"remaining":5}`),
				)
			})
		})

		Context("When deck created with invalid cards option", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{id: "ABC-DEF"}
			router.POST("/deck", deck.CreateBuilder(deckRepo))

			req, _ := http.NewRequest("POST", "/deck?cards=AS,KS,ZD,2C,KH", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 400", func() {
				Expect(rr.Code).To(Equal(400))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"error":"invalid card code provided"}`),
				)
			})
		})
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deck Handlers Suite")
}
