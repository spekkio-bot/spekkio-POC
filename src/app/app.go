package app

import (
	"fmt"
	"net/http"
)

type App struct {

}

func (a *App) Run() {
	fmt.Printf("Serving on 127.0.0.1:2000.\n")
	http.ListenAndServe("127.0.0.1:2000", nil)
}
