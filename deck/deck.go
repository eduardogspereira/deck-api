package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eduardogspereira/deck-api/card"
)

// Options implements the options object that
// can be provided for a new deck.
type Options struct {
	Shuffle     bool
	WantedCards []string
}

// Deck implements the deck object.
type Deck struct {
	ID       string
	Shuffled bool
	Cards    []card.Card
}

// Remaining returns how many cards are in the deck.
func (d Deck) Remaining() int {
	return len(d.Cards)
}

// Shuffle implements the card shuffle on the deck
func (d Deck) Shuffle() {
	d.Shuffled = true

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// New returns a new deck based on the options provided.
func New(options Options) (Deck, error) {
	var deck Deck
	var err error

	cardCodes := makeAllCardCodes()
	if len(options.WantedCards) > 0 {
		cardCodes = options.WantedCards
	}

	for _, cardCode := range cardCodes {
		c, err := card.FromCode(cardCode)
		if err != nil {
			return deck, fmt.Errorf("failed to create new deck: %v", err)
		}
		deck.Cards = append(deck.Cards, c)
	}

	if options.Shuffle {
		deck.Shuffle()
	}

	return deck, err
}

func makeAllCardCodes() []string {
	var cardCodes []string

	suits := []string{"S", "D", "C", "H"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			cardCodes = append(cardCodes, value+suit)
		}
	}

	return cardCodes
}
