package server

import (
	"github.com/davyzhang/agw"
	"github.com/spekkio-bot/spekkio/src/app"
)

func Run() {
	app := &app.App{
		Config: &app.AppConfig{},
	}
	app.Config.Load()
	app.Initialize()
	app.Run()
}

func RunOnLambda() {
	app := &app.App{
		Config: &app.AppConfig{},
	}
	app.Config.Load()
	app.Initialize()
	lambda.Start(agw.Handler(app.Router))
}
