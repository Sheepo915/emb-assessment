package main

import (
	"sheepo.com/emb_assessment/internal/app"
	"sheepo.com/emb_assessment/internal/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg := config.NewConfig()
	cfg.ParseFlag()

	app := app.NewApp(cfg)

	app.Run()
}
