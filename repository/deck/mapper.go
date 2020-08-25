package deck

import (
	"github.com/eduardogspereira/deck-api/domains/deck"
	"github.com/gofrs/uuid"
)

func toDBModel(entity *deck.Deck) *Deck {
	var d Deck

	d.Shuffled = entity.Shuffled
	d.CardCodes = entity.CardsToCodes()

	if value, err := uuid.FromString(entity.ID); err == nil {
		d.ID = value
	}

	return &d
}

func toDomainModel(entity *Deck) *deck.Deck {
	var d deck.Deck

	d.ID = entity.ID.String()
	d.Shuffled = entity.Shuffled
	d.Cards = d.CodesToCards(entity.CardCodes)

	return &d
}
