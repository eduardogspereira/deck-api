package deck

import (
	domain "github.com/eduardogspereira/deck-api/domains/deck"
	"github.com/jinzhu/gorm"
)

// Repo contains the repository implementation for the deck
type Repo struct {
	db *gorm.DB
}

// New creates a new deck repository.
func New(db *gorm.DB) *Repo {
	db.AutoMigrate(&Deck{})

	return &Repo{
		db: db,
	}
}

// Save the deck in the database.
func (r *Repo) Save(deck *domain.Deck) (*domain.Deck, error) {
	entity := toDBModel(deck)

	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return toDomainModel(entity), nil
}

// FindByID finds the deck by its ID.
func (r *Repo) FindByID(deckID string) (*domain.Deck, error) {
	result := &Deck{}
	query := r.db.Where("id = ?", deckID).First(result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return toDomainModel(result), nil
}

// Update the deck.
func (r *Repo) Update(deck *domain.Deck) error {
	entity := toDBModel(deck)
	query := r.db.Model(&entity).Where("id = ?", entity.ID).Update(entity)
	return query.Error
}
