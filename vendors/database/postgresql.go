package database

import (
	"fmt"

	"github.com/eduardogspereira/deck-api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // It's required to import the postgresql driver
)

// Connect to a database handle from a connection string.
func Connect(config config.Database) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", config.Host, config.Port, config.DB, config.User, config.Password)
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}
