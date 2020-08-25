package deck

import (
	"fmt"
	"net/http"

	"github.com/eduardogspereira/deck-api/domains/deck"
	deckRepo "github.com/eduardogspereira/deck-api/repository/deck"
	"github.com/gin-gonic/gin"
)

// CreateBuilder handles the creation of a new deck.
func CreateBuilder(deckRepo deckRepo.Repository) func(c *gin.Context) {
	create := func(c *gin.Context) {
		options, err := bindCreateOptions(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		deckOptions := deck.Options{
			Shuffle:     options.Shuffle,
			WantedCards: options.Cards,
		}
		d, err := deck.New(deckOptions)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		d, err = deckRepo.Save(d)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusCreated, toCreateResponse(d))
	}

	return create
}
