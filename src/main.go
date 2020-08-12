package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spekkio-bot/spekkio/src/server"
)

func LoadFromDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("err: no .env file found\n")
	}
}

func InvalidArgs() {
	log.Fatal("err: invalid args\n")
}

func main() {
	LoadFromDotenv()
	args := os.Args
	if len(args) == 1 {
		log.Printf("no args selected - no action will be taken.\n")
	} else if len(args) == 2 {
		switch args[1] {
		case "dev":
			server.Start()
		default:
			InvalidArgs()
		}
	} else {
		InvalidArgs()
	}
}
