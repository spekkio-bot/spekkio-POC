package app

import (
	"fmt"
	"net/http"
)

type App struct {
	Config *AppConfig
}

func (a *App) Run() {
	fmt.Printf("Serving on %s.\n", a.Config.Server.GetAddr())
	http.ListenAndServe(a.Config.Server.GetAddr(), nil)
}
