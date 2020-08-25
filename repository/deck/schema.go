package deck

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Deck defines the database schema for a deck
type Deck struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Shuffled  bool           `gorm:"default:false"`
	CardCodes pq.StringArray `gorm:"type:varchar(3)[]"`
}

// BeforeCreate is a Hook function to set the UUID
// before the deck goes into the database.
func (base *Deck) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}
