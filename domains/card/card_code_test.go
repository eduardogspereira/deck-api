package card_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/eduardogspereira/deck-api/domains/card"
)

var _ = Describe("CardCode", func() {
	Describe("FromCode", func() {
		Context("When card code is valid", func() {
			It("should create a Card with the correct properties for QH", func() {
				c, _ := card.FromCode("QH")
				Expect(c.Code).To(Equal("QH"))
				Expect(c.Value).To(Equal("QUEEN"))
				Expect(c.Suit).To(Equal("HEARTS"))
			})

			It("should create a Card with the correct properties for 4D", func() {
				c, _ := card.FromCode("4D")
				Expect(c.Code).To(Equal("4D"))
				Expect(c.Value).To(Equal("4"))
				Expect(c.Suit).To(Equal("DIAMONDS"))
			})

			It("should create a Card with the correct properties for 10S", func() {
				c, _ := card.FromCode("10S")
				Expect(c.Code).To(Equal("10S"))
				Expect(c.Value).To(Equal("10"))
				Expect(c.Suit).To(Equal("SPADES"))
			})
		})

		Context("When card code is invalid", func() {
			It("should return an error for EZ", func() {
				_, err := card.FromCode("EZ")
				Expect(err).To(Not(BeNil()))
			})

			It("should return an error for an empty string", func() {
				_, err := card.FromCode("")
				Expect(err).To(Not(BeNil()))
			})
		})
	})
})

func TestCardCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CardCode Suite")
}
