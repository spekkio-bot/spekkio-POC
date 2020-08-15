package app

import (
	//"context"
	"database/sql"
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davyzhang/agw"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	"github.com/spekkio-bot/spekkio/src/app/controller"
)

type App struct {
	Config  *AppConfig
	Db      *sql.DB
	Router  *mux.Router
	Handler http.Handler
}

func (a *App) Run() {
	switch (a.Config.Platform) {
	case "default":
		fmt.Printf("serving on %s.\n", a.Config.Server.GetAddr())
		http.ListenAndServe(a.Config.Server.GetAddr(), a.Handler)
	case "lambda":
		fmt.Printf("running on aws lambda mode.\n")
		lambda.Start(agw.Handler(a.Handler))
		/*
		lambda.Start(func() agw.GatewayHandler {
			return func(ctx context.Context, event json.RawMessage) (interface{}, error) {
				agp := agw.NewAPIGateParser(event)
				return agw.Process(agp, a.Handler), nil
			}
		}())
		*/
	default:
		log.Fatal("err: invalid platform option.\n")
	}
}

func (a *App) Initialize() {
	a.ConnectToDb()
	a.Router = mux.NewRouter()
	a.SetRoutes()
	originsOk := handlers.AllowedOrigins([]string{a.Config.AllowedOrigins})
	a.Handler = alice.New(handlers.CORS(originsOk)).Then(handlers.CombinedLoggingHandler(os.Stdout, a.Router))
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
	controller.Ping(w, r)
}
