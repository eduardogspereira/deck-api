package deck_test

import (
	"errors"

	"github.com/eduardogspereira/deck-api/domains/deck"
)

type MockedDeckRepo struct {
	id                string
	returnErrorOnSave bool
}

func (m MockedDeckRepo) Save(d deck.Deck) (deck.Deck, error) {
	d.ID = m.id

	if m.returnErrorOnSave {
		return d, errors.New("error on save")
	}

	return d, nil
}
