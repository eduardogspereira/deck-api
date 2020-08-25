package deck

import (
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

// LoadBuilder handles the load of a deck.
func LoadBuilder(deckRepo deckRepo.Repository) func(c *gin.Context) {
	load := func(c *gin.Context) {
		deckID := c.Param("deckID")

		d, err := deckRepo.FindByID(deckID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		if d.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "no deck found for the specified id"})
			return
		}

		c.JSON(http.StatusOK, toLoadReponse(d))
	}

	return load
}

// DrawCardBuilder draws cards from the specified deck.
func DrawCardBuilder(deckRepo deckRepo.Repository) func(c *gin.Context) {
	drawCard := func(c *gin.Context) {
		options, err := bindDrawCardOptions(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid count parameter provided"})
			return
		}

		deckID := c.Param("deckID")

		d, err := deckRepo.FindByID(deckID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		if d.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "no deck found for the specified id"})
			return
		}

		if d.Remaining() < options.Count {
			c.JSON(http.StatusUnprocessableEntity,
				gin.H{
					"message":   "the count provided is greater than the remaining cards in the deck",
					"remaining": d.Remaining(),
				})
			return
		}

		cards := d.DrawCards(options.Count)

		err = deckRepo.Update(d)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		c.JSON(http.StatusOK, toDrawCardResponse(cards))
	}

	return drawCard
}
