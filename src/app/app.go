package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Config *AppConfig
	Db     *sql.DB
	Router *mux.Router
}

func (a *App) Run() {
	fmt.Printf("Serving on %s.\n", a.Config.Server.GetAddr())
	http.ListenAndServe(a.Config.Server.GetAddr(), nil)
}

func (a *App) Initialize() {
	a.ConnectToDb()
	a.Router = mux.NewRouter()
	a.SetRoutes()
}

func (a *App) ConnectToDb() {
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

func (a *App) SetRoutes() {
	a.Get("/", a.Ping)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Ping(w http.ResponseWriter, r *http.Request) {

}
