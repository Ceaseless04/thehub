package config

import (
	"log"
	"os"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port string;
	dbURL string;
	dbKEY string;
}

func LoadConfig() (*config, error) {
	port := os.Getenv("PORT");
	if port == "" {
		log.Fatal("Server port is required");
	}

	dbURL := os.Getenv("SUPA_URL");
	if dbURL == "" {
		log.Fatal("The URL for the Database is required");
	}

	dbKEY := os.Getenv("SUPA_KEY");
	if dbKEY == "" {
		log.Fatal("The KEY for the Database is required");
	}

	return &Config {
		Port: port,
		dbURL: dbURL,
		dbKEY: dbKEY
	}, nil
}
