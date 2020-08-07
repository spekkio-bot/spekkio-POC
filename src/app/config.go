package app

import (
	"fmt"
	"os"
)

type AppConfig struct {
	Database       *DatabaseConfig
	Server         *ServerConfig
	AllowedOrigins string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Schema   string
	SslMode  string
}

type ServerConfig struct {
	Host string
	Port string
}

func (ac *AppConfig) Load() {
	ac.Server = &ServerConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	ac.Database = &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Schema:   os.Getenv("DB_SCHEMA"),
		SslMode:  os.Getenv("DB_SSLMODE"),
	}

	ac.AllowedOrigins = os.Getenv("ORIGINS_ALLOWED")

	if len(ac.Server.Host) == 0 {
		ac.Server.Host = "127.0.0.1"
	}

	if len(ac.Server.Port) == 0 {
		ac.Server.Port = "2000"
	}

	if len(ac.AllowedOrigins) == 0 {
		ac.AllowedOrigins = ""
	}

	if len(ac.Database.SslMode) == 0 {
		ac.Database.SslMode = "prefer"
	}
}

func (sc *ServerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", sc.Host, sc.Port)
}

func (dc *DatabaseConfig) GetInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dc.Host,
		dc.Port,
		dc.User,
		dc.Password,
		dc.Schema,
		dc.SslMode)
}
