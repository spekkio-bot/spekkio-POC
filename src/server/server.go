package server

import (
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
