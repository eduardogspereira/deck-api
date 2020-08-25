package deck

import (
	"github.com/eduardogspereira/deck-api/domains/card"
	"github.com/eduardogspereira/deck-api/domains/deck"
)

func toCreateResponse(entity *deck.Deck) createResponse {
	return createResponse{
		DeckID:    entity.ID,
		Shuffled:  entity.Shuffled,
		Remaining: entity.Remaining(),
	}
}

func toCardResponse(cardEntity card.Card) cardResponse {
	return cardResponse{
		Value: cardEntity.Value,
		Suit:  cardEntity.Suit,
		Code:  cardEntity.Code,
	}
}

func toLoadReponse(entity *deck.Deck) loadResponse {
	response := loadResponse{
		DeckID:    entity.ID,
		Shuffled:  entity.Shuffled,
		Remaining: entity.Remaining(),
	}

	response.Cards = make([]cardResponse, 0)
	for _, c := range entity.Cards {
		response.Cards = append(response.Cards, toCardResponse(c))
	}

	return response
}
