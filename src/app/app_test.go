package app

import (
	"os"
	"testing"
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
		t.Errorf("GetAddr() returned unexpected server address\n\ngot %v\nexpected %v\n", db, wantDb)
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
		t.Errorf("GetAddr() returned unexpected server address\n\ngot %v\nexpected %v\n", db, wantDb)
	}
}
