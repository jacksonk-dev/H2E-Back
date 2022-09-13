﻿# Entertainment Back-end
Backend service that fetches data from external APIs such as TMDB and Trakt.<br />
Requires Go, preferably the latest version.
## Set up for development
Create an `app.dev.yaml` file in the root directory and add the following:<br />
env_variables:<br />
  PORT: 8082<br />
  CLIENT_ORIGIN: http://localhost:4200<br />
  TMDB_API_KEY: <YOUR_TMDB_API_KEY><br />
  TRAKT_API_KEY: <YOUR_TRAKT_API_KEY><br />

## Install dependencies

Run `go build`. Checkout https://stackoverflow.com/questions/69474609/installing-all-dependencies-from-a-go-mod-file

## Run app
Run `go run main.go` or `go run .`
