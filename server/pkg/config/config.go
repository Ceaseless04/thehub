package config

import (
	"log"
	"os"
)

type config struct {
	Port string;
	dbURL string;
	dbKEY string;
}

func LoadConfig() (*config, error) {
	
}
