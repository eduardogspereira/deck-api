package config

import "github.com/gobuffalo/envy"

// Config contains all the required config variables
// required for the server execution.
type Config struct {
	Database
}

// Database contains the credentials for accessing the DB.
type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

// New creates a Config structure from the environment variables.
func New() (Config, error) {
	var config Config
	var err error

	err = envy.Load()
	if err != nil {
		return config, err
	}

	config = Config{
		Database: Database{
			Host:     mustGetEnv("PGHOST"),
			Port:     mustGetEnv("PGPORT"),
			User:     mustGetEnv("PGUSER"),
			DB:       mustGetEnv("PGDATABASE"),
			Password: mustGetEnv("PGPASSWORD"),
		},
	}

	return config, err
}

func mustGetEnv(name string) (value string) {
	value, err := envy.MustGet(name)
	if err != nil {
		panic("The required environment variable " + name + " was not found.")
	}
	return
}
