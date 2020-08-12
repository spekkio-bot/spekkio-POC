package server

import (
	"github.com/spekkio-bot/spekkio/src/app"
)

func Start() {
	app := &app.App{
		Config: &app.AppConfig{},
	}
	app.Config.Load()
	app.Initialize()
	app.Run()
}
