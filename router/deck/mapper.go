package deck

import "github.com/eduardogspereira/deck-api/domains/deck"

func toCreateResponse(entity deck.Deck) CreateResponse {
	return CreateResponse{
		DeckID:    entity.ID,
		Shuffled:  entity.Shuffled,
		Remaining: entity.Remaining(),
	}
}
