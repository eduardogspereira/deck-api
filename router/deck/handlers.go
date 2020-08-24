package deck

import (
	"encoding/json"
	"net/http"

	"github.com/eduardogspereira/deck-api/domains/deck"
	deckRepo "github.com/eduardogspereira/deck-api/repository/deck"
)

// CreateBuilder handles the creation of a new deck.
func CreateBuilder(deckRepo deckRepo.Repository) func(w http.ResponseWriter, r *http.Request) {
	create := func(w http.ResponseWriter, r *http.Request) {
		options := deck.Options{}
		d, err := deck.New(options)
		if err != nil {
			panic(err)
		}

		d, err = deckRepo.Save(d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(toCreateResponse(d))
	}

	return create
}
