package deck

import "github.com/eduardogspereira/deck-api/domains/deck"

// Repository interface defines the methods for the deck
//  repository.
type Repository interface {
	Save(*deck.Deck) (*deck.Deck, error)
}
