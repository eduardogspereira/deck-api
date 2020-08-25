package deck

import "github.com/eduardogspereira/deck-api/domains/deck"

// Repository interface defines the methods for the deck
//  repository.
type Repository interface {
	FindById(string) (*deck.Deck, error)
	Save(*deck.Deck) (*deck.Deck, error)
}
