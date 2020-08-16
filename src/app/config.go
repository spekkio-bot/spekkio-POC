package app

import (
	"fmt"
	"os"
)

// AppConfig defines the settings of the app that can be configured.
type AppConfig struct {
	Database       *DatabaseConfig
	Server         *ServerConfig
	AllowedOrigins string
	Platform       string
}

// DatabaseConfig defines the parameters needed to connect to a database.
// Currently, only Postgres databases are supported.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Schema   string
	SslMode  string
}

// ServerConfig defines the parameters needed to serve the app on an IPv4 interface.
type ServerConfig struct {
	Host string
	Port string
}

// Load will load environmental variables into AppConfig and set default values where applicable.
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
	ac.Platform = os.Getenv("PLATFORM")

	if len(ac.Server.Host) == 0 {
		ac.Server.Host = "127.0.0.1"
	}

	if len(ac.Server.Port) == 0 {
		ac.Server.Port = "2000"
	}

	if len(ac.AllowedOrigins) == 0 {
		ac.AllowedOrigins = ""
	}

	if len(ac.Platform) == 0 {
		ac.Platform = "default"
	}

	if len(ac.Database.SslMode) == 0 {
		ac.Database.SslMode = "prefer"
	}
}

// GetAddr returns the full IPv4 interface defined in ServerConfig.
func (sc *ServerConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", sc.Host, sc.Port)
}

// GetInfo returns the parameters defined in DatabaseConfig.
func (dc *DatabaseConfig) GetInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dc.Host,
		dc.Port,
		dc.User,
		dc.Password,
		dc.Schema,
		dc.SslMode)
}
