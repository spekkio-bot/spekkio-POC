package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spekkio-bot/spekkio/src/app"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		serve()
	} else if len(args) == 2 {
		switch args[1] {
		case "dev":
			loadFromDotenv()
			serve()
		default:
			invalidArgs()
		}
	} else {
		invalidArgs()
	}
}

func serve() {
	app := &app.App{
		Config: &app.AppConfig{},
	}
	app.Config.Load()
	app.Initialize()
	app.Run()
}

func loadFromDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("err: no .env file found\n")
	}
}

func invalidArgs() {
	log.Fatal("err: invalid args\n")
}
