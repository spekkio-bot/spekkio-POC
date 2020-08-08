package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type App struct {
	Config *AppConfig
	Db     *sql.DB
}

func (a *App) Run() {
	fmt.Printf("Serving on %s.\n", a.Config.Server.GetAddr())
	http.ListenAndServe(a.Config.Server.GetAddr(), nil)
}

func (a *App) Initialize() {
	var err error

	a.Db, err = sql.Open("postgres", a.Config.Database.GetInfo())
	if err != nil {
		fmt.Printf("fatal err: cannot connect to database.\n")
		log.Fatal(err)
	} else {
		fmt.Printf("successfully connected to database!\n")
	}

	err = a.Db.Ping()
	if err != nil {
		fmt.Printf("fatal err: cannot ping database.\n")
		log.Fatal(err)
	} else {
		fmt.Printf("successfully pinged database!\n")
	}
}
