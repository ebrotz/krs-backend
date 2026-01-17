package config

import "os"

// Config represents top level application configuration. Items that go into
// this struct are typically items that are set via environment variables.
type Config struct {
	DatabaseUrl string
}

func ConfigFromEnv() Config {
	databaseUrl, _ := os.LookupEnv("DATABASE_URL")

	return Config{
		DatabaseUrl: databaseUrl,
	}
}
