package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/park-junha/agw"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres database connection requires postgres driver
	"github.com/spekkio-bot/spekkio/src/app/controller"
)

// App defines the various components of the app.
type App struct {
	Config *AppConfig
	Db     *sql.DB
	Router *mux.Router
}

// Run will run the app depending on what platform it is configured to run on.
func (a *App) Run() {
	switch a.Config.Platform {
	case "default":
		fmt.Printf("serving on %s.\n", a.Config.Server.GetAddr())
		log.Fatal(http.ListenAndServe(a.Config.Server.GetAddr(), a.Router))
	case "lambda":
		fmt.Printf("running on aws lambda mode.\n")
		lambda.Start(agw.Handler(a.Router))
	default:
		log.Fatal("err: invalid platform option.\n")
	}
}

// Initialize will create the various components of the app needed to run it from a.Config.
// The app can be configured with a.Config.Load().
func (a *App) Initialize() {
	a.ConnectToDb()
	a.Router = mux.NewRouter()
	a.SetRoutes()
	a.SetMiddleware()
}

// ConnectToDb will connect the app to the database specified by a.Config.
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

// SetRoutes will initialize all the routes of the app.
func (a *App) SetRoutes() {
	a.Get("/", a.Ping)
	a.Post("/scrumify", a.Scrumify)
	a.Router.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}

// SetMiddleware will initialize middleware handlers for the router.
func (a *App) SetMiddleware() {
	a.Router.Use(logger)
}

// Get is a wrapper for resources requested with GET.
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post is a wrapper for resources requested with POST.
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Ping calls the Ping controller.
func (a *App) Ping(w http.ResponseWriter, r *http.Request) {
	controller.Ping(w, r)
}

// Scrumify calls the Scrumify controller.
func (a *App) Scrumify(w http.ResponseWriter, r *http.Request) {
	controller.Scrumify(a.Db, w, r)
}
