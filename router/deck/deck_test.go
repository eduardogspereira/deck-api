package deck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardogspereira/deck-api/domains/card"
	deckDomain "github.com/eduardogspereira/deck-api/domains/deck"
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

	Describe("Load", func() {
		Context("When requests a deck", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{
				deck: deckDomain.Deck{
					ID:       "DECK-ABC",
					Shuffled: false,
					Cards: []card.Card{
						card.Card{Code: "1A", Suit: "ACES", Value: "1"},
						card.Card{Code: "10C", Suit: "CLUBS", Value: "10"},
					},
				},
			}
			router.GET("/deck/:deckId", deck.LoadBuilder(deckRepo))

			req, _ := http.NewRequest("GET", "/deck/123", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 200", func() {
				Expect(rr.Code).To(Equal(200))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(
						`{"deck_id":"DECK-ABC","shuffled":false,"remaining":2,"cards":[{"value":"1",` +
							`"suit":"ACES","code":"1A"},{"value":"10","suit":"CLUBS","code":"10C"}]}`,
					),
				)
			})
		})

		Context("When requests a deck with no remaining cards", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{
				deck: deckDomain.Deck{
					ID:       "DECK-ABC",
					Shuffled: false,
				},
			}
			router.GET("/deck/:deckId", deck.LoadBuilder(deckRepo))

			req, _ := http.NewRequest("GET", "/deck/123", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 200", func() {
				Expect(rr.Code).To(Equal(200))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(
						`{"deck_id":"DECK-ABC","shuffled":false,"remaining":0,"cards":[]}`,
					),
				)
			})
		})

		Context("When deckRepo.FindById returns an error", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{returnErrorOnFindById: true}
			router.GET("/deck/:deckID", deck.LoadBuilder(deckRepo))

			req, _ := http.NewRequest("GET", "/deck/123", nil)
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

		Context("When no deck found within the specified ID", func() {
			router := routerFactory()
			rr := httptest.NewRecorder()

			deckRepo := MockedDeckRepo{}
			router.GET("/deck/:deckID", deck.LoadBuilder(deckRepo))

			req, _ := http.NewRequest("GET", "/deck/123", nil)
			router.ServeHTTP(rr, req)

			It("should have the status_code equal 404", func() {
				Expect(rr.Code).To(Equal(404))
			})

			It("should return the correct body", func() {
				Expect(rr.Body.String()).To(
					Equal(`{"message":"no deck found for the specified id"}`),
				)
			})
		})
	})
})

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deck Handlers Suite")
}
