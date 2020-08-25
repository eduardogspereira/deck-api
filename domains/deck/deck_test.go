package deck_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/eduardogspereira/deck-api/domains/deck"
)

var _ = Describe("Deck", func() {
	Describe("New", func() {
		Context("When default options are used", func() {
			options := deck.Options{}
			d, _ := deck.New(options)

			It("should have shuffled property equal false", func() {
				Expect(d.Shuffled).To(BeFalse())
			})

			It("should should contains 52 remaining cards", func() {
				Expect(d.Remaining()).To(Equal(52))
			})

			It("first and last cards should be in order", func() {
				Expect(d.Cards[0].Code).To(Equal("AS"))
				Expect(d.Cards[51].Code).To(Equal("KH"))
			})
		})

		Context("When shuffle option is used", func() {
			options := deck.Options{
				Shuffle: true,
			}
			d, _ := deck.New(options)

			It("should have shuffled property equal true", func() {
				Expect(d.Shuffled).To(BeTrue())
			})

			It("first 5 cards should not be in order", func() {
				Expect(d.Cards[0].Code == "AS" &&
					d.Cards[1].Code == "2S" &&
					d.Cards[2].Code == "3S" &&
					d.Cards[3].Code == "4S" &&
					d.Cards[4].Code == "5S").To(BeFalse())
			})
		})

		Context("When wanted cards option is used", func() {
			options := deck.Options{
				WantedCards: []string{"AS", "KD", "AC", "2C", "KH"},
			}
			d, _ := deck.New(options)

			It("should have the correct remaining cards", func() {
				Expect(d.Remaining()).To(Equal(5))
			})
		})

		Context("When invalid option is provided", func() {
			options := deck.Options{
				WantedCards: []string{"ZD"},
			}
			_, err := deck.New(options)

			It("should return an error", func() {
				Expect(err).To(Not(BeNil()))
			})
		})
	})
})

func TestDeck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deck Suite")
}
