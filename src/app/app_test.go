package app

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
)

func setEnvDefaultVars() {
	os.Setenv("HOST", "")
	os.Setenv("PORT", "")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "Password123!")
	os.Setenv("DB_SCHEMA", "SourceDb")
	os.Setenv("DB_SSLMODE", "")
	os.Setenv("ORIGINS_ALLOWED", "")
	os.Setenv("PLATFORM", "")
}

func setEnvVars() {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "Password123!")
	os.Setenv("DB_SCHEMA", "SourceDb")
	os.Setenv("DB_SSLMODE", "require")
	os.Setenv("ORIGINS_ALLOWED", "*")
	os.Setenv("PLATFORM", "default")
}

func (a *App) runTestRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func TestAppConfigLoad(t *testing.T) {
	app := &App{
		Config: &AppConfig{},
	}

	setEnvVars()
	app.Config.Load()

	if app.Config.Server.Host != "localhost" {
		t.Errorf("AppConfig loaded unexpected Server.Host value\n\ngot %v\nexpected %v\n", app.Config.Server.Host, "localhost")
	}
	if app.Config.Server.Port != "8080" {
		t.Errorf("AppConfig loaded unexpected Server.Port value\n\ngot %v\nexpected %v\n", app.Config.Server.Port, "8080")
	}
	if app.Config.Database.Host != "127.0.0.1" {
		t.Errorf("AppConfig loaded unexpected Database.Host value\n\ngot %v\nexpected %v\n", app.Config.Database.Host, "127.0.0.1")
	}
	if app.Config.Database.Port != "5432" {
		t.Errorf("AppConfig loaded unexpected Database.Port value\n\ngot %v\nexpected %v\n", app.Config.Database.Port, "5432")
	}
	if app.Config.Database.User != "admin" {
		t.Errorf("AppConfig loaded unexpected Database.User value\n\ngot %v\nexpected %v\n", app.Config.Database.User, "admin")
	}
	if app.Config.Database.Password != "Password123!" {
		t.Errorf("AppConfig loaded unexpected Database.Password value\n\ngot %v\nexpected %v\n", app.Config.Database.Password, "Password123!")
	}
	if app.Config.Database.Schema != "SourceDb" {
		t.Errorf("AppConfig loaded unexpected Database.Schema value\n\ngot %v\nexpected %v\n", app.Config.Database.Schema, "SourceDb")
	}
	if app.Config.Database.SslMode != "require" {
		t.Errorf("AppConfig loaded unexpected Database.SslMode value\n\ngot %v\nexpected %v\n", app.Config.Database.SslMode, "require")
	}
	if app.Config.AllowedOrigins != "*" {
		t.Errorf("AppConfig loaded unexpected AllowedOrigins value\n\ngot %v\nexpected %v\n", app.Config.AllowedOrigins, "*")
	}
	if app.Config.Platform != "default" {
		t.Errorf("AppConfig loaded unexpected Platform value\n\ngot %v\nexpected %v\n", app.Config.Platform, "default")
	}

	ip := app.Config.Server.GetAddr()
	db := app.Config.Database.GetInfo()
	wantIp := "localhost:8080"
	wantDb := "host=127.0.0.1 port=5432 user=admin password=Password123! dbname=SourceDb sslmode=require"

	if ip != wantIp {
		t.Errorf("GetAddr() returned unexpected server address\n\ngot %v\nexpected %v\n", ip, wantIp)
	}
	if db != wantDb {
		t.Errorf("GetInfo() returned unexpected database info\n\ngot %v\nexpected %v\n", db, wantDb)
	}
}

func TestAppConfigLoadDefaults(t *testing.T) {
	app := &App{
		Config: &AppConfig{},
	}

	setEnvDefaultVars()
	app.Config.Load()

	if app.Config.Server.Host != "127.0.0.1" {
		t.Errorf("AppConfig loaded unexpected Server.Host value\n\ngot %v\nexpected %v\n", app.Config.Server.Host, "127.0.0.1")
	}
	if app.Config.Server.Port != "2000" {
		t.Errorf("AppConfig loaded unexpected Server.Port value\n\ngot %v\nexpected %v\n", app.Config.Server.Port, "2000")
	}
	if app.Config.Database.Host != "127.0.0.1" {
		t.Errorf("AppConfig loaded unexpected Database.Host value\n\ngot %v\nexpected %v\n", app.Config.Database.Host, "127.0.0.1")
	}
	if app.Config.Database.Port != "5432" {
		t.Errorf("AppConfig loaded unexpected Database.Port value\n\ngot %v\nexpected %v\n", app.Config.Database.Port, "5432")
	}
	if app.Config.Database.User != "admin" {
		t.Errorf("AppConfig loaded unexpected Database.User value\n\ngot %v\nexpected %v\n", app.Config.Database.User, "admin")
	}
	if app.Config.Database.Password != "Password123!" {
		t.Errorf("AppConfig loaded unexpected Database.Password value\n\ngot %v\nexpected %v\n", app.Config.Database.Password, "Password123!")
	}
	if app.Config.Database.Schema != "SourceDb" {
		t.Errorf("AppConfig loaded unexpected Database.Schema value\n\ngot %v\nexpected %v\n", app.Config.Database.Schema, "SourceDb")
	}
	if app.Config.Database.SslMode != "prefer" {
		t.Errorf("AppConfig loaded unexpected Database.SslMode value\n\ngot %v\nexpected %v\n", app.Config.Database.SslMode, "prefer")
	}
	if app.Config.AllowedOrigins != "" {
		t.Errorf("AppConfig loaded unexpected AllowedOrigins value\n\ngot %v\nexpected %v\n", app.Config.AllowedOrigins, "")
	}
	if app.Config.Platform != "default" {
		t.Errorf("AppConfig loaded unexpected Platform value\n\ngot %v\nexpected %v\n", app.Config.Platform, "default")
	}

	ip := app.Config.Server.GetAddr()
	db := app.Config.Database.GetInfo()
	wantIp := "127.0.0.1:2000"
	wantDb := "host=127.0.0.1 port=5432 user=admin password=Password123! dbname=SourceDb sslmode=prefer"

	if ip != wantIp {
		t.Errorf("GetAddr() returned unexpected server address\n\ngot %v\nexpected %v\n", ip, wantIp)
	}
	if db != wantDb {
		t.Errorf("GetInfo() returned unexpected database info\n\ngot %v\nexpected %v\n", db, wantDb)
	}
}

func TestAppRouter(t *testing.T) {
	app := &App{
		Config: &AppConfig{},
	}

	setEnvVars()
	app.Config.Load()

	app.Router = mux.NewRouter()
	app.SetRoutes()
	if app.Router.Get("Ping") == nil {
		t.Errorf("Router did not register Ping route\n")
	}
	if app.Router.Get("Scrumify") == nil {
		t.Errorf("Router did not register Scrumify route\n")
	}
	if app.Router.NotFoundHandler == nil {
		t.Errorf("Router did not register NotFoundHandler\n")
	}

	var err error
	//var mock sqlmock.Sqlmock
	app.Db, _, err = sqlmock.New()
	if err != nil {
		t.Fatalf("encountered error %s while initializing mock database\n", err)
	}
	defer app.Db.Close()
	app.SetMiddleware()

	reqPing, _ := http.NewRequest("GET", "/", nil)
	reqNotFound, _ := http.NewRequest("GET", "/asdfghjkl", nil)

	respPing := app.runTestRequest(reqPing)
	respNotFound := app.runTestRequest(reqNotFound)

	if statusPing := respPing.Code; statusPing != 200 {
		t.Errorf("Ping returned wrong status code:\ngot %v\nwant %v\n", statusPing, 200)
	}

	wantPing := `{"message":"Request successful."}`
	gotPing := respPing.Body.String()
	if gotPing != wantPing {
		t.Errorf("Ping return unexpected body:\ngot %v\nwant %v\n", gotPing, wantPing)
	}

	if statusNotFound := respNotFound.Code; statusNotFound != 404 {
		t.Errorf("NotFound returned wrong status code:\ngot %v\nwant %v\n", statusNotFound, 404)
	}

	wantNotFound := `{"message":"What do you want?","error":"resource not found."}`
	gotNotFound := respNotFound.Body.String()
	if gotNotFound != wantNotFound {
		t.Errorf("NotFound return unexpected body:\ngot %v\nwant %v\n", gotNotFound, wantNotFound)
	}
}
