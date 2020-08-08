package main

import (
	"log"

	"github.com/spekkio-bot/spekkio/src/app"
	"github.com/joho/godotenv"
)

func LoadFromDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("err: no .env file found\n")
	}
}

func main() {
	LoadFromDotenv()
	app := &app.App{
		Config: &app.AppConfig{},
	}
	app.Config.Load()
	app.Initialize()
	app.Run()
}
