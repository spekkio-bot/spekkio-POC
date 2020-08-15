# Spekkio - Source Code

This directory contains source code for the main Spekkio app.

## Contents

```
├── app                             # App directory
│   ├── controller                  # App controller
│       ├── common.go               # Common controller functions
│       ├── index.go                # Index route
│   ├── model                       # App models
│       ├── models.go               # Models for app data
│   ├── app.go                      # Top-level app source code
│   ├── config.go                   # Functions to configure the app
├── .env                            # Environmental variables
├── main.go                         # Main script
```

## First Time Setup

1. Create a `.env` file. You may copy the `.env.example` file by running:
```
cp .env.example .env
```

2. Configure the following settings in your `.env` file:
```
# Database credentials (currently support Postgres only)
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=samplepassword
DB_SCHEMA=sampledb
DB_SSLMODE=prefer             # if undefined, app defaults to prefer

# Dev server configuration
HOST=127.0.0.1                # if undefined, app defaults to 127.0.0.1
PORT=2000                     # if undefined, app defaults to 2000
PLATFORM=default              # if undefined, app defaults to default

# OPTIONAL: Allowed origins for CORS
ORIGINS_ALLOWED=*             # if undefined, app defaults to no origins
```

## Run the App

### Build and run

1. Run `go build main.go` to build your app.
2. Run `./main dev` to execute your compiled program.

### Run without build

1. Run `go run main.go dev` to run your app without building it.
