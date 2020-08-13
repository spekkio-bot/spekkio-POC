package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
	"github.com/spekkio-bot/spekkio/src/server"
	"github.com/spekkio-bot/spekkio/src/serverless"
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
	args := os.Args
	if len(args) == 1 {
		LoadFromDotenv()
		server.Run()
	} else if len(args) == 2 {
		switch args[1] {
		case "lambda":
			server.RunOnLambda()
		default:
			InvalidArgs()
		}
	} else {
		InvalidArgs()
	}
}
