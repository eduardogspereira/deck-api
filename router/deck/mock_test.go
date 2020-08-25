package deck_test

import (
	"errors"

	"github.com/eduardogspereira/deck-api/domains/deck"
)

type MockedDeckRepo struct {
	id                    string
	returnErrorOnSave     bool
	deck                  deck.Deck
	returnErrorOnFindById bool
}

func (m MockedDeckRepo) Save(d *deck.Deck) (*deck.Deck, error) {
	d.ID = m.id

	if m.returnErrorOnSave {
		return d, errors.New("error on save")
	}

	return d, nil
}

func (m MockedDeckRepo) FindByID(deckId string) (*deck.Deck, error) {
	if m.returnErrorOnFindById {
		return &m.deck, errors.New("error on find by id")
	}

	return &m.deck, nil
}

func (m MockedDeckRepo) Update(d *deck.Deck) error {
	return nil
}
