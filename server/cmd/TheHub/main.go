package main

import (
	"log"
	"github.com/Ceaseless04/TheHub/internal/app"
	"github.com/Ceaseless04/TheHub/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig();
	if err != nil {
		log.Fatalf("Could not load config: %v", err);
	}

	app, err := app.NewApp(cfg);
	if err != nil {
		log.Fatalf("Could not run app: %v", err);
	}
}
