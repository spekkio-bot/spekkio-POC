package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spekkio-bot/spekkio/src/app"
)

func LoadFromDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("err: no .env file found\n")
	}
}

func main() {
	LoadFromDotenv()
	args := os.Args()

	app := &app.App{
		Config: &app.AppConfig{},
	}
	app.Config.Load()
	app.Initialize()
	app.Run()
}
