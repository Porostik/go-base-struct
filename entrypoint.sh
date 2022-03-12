#!/bin/sh

source .env

if [[ "$ENVIRONMENT_MODE" == "production" ]]
then
    go build -o /docker-main-app ./cmd/main
    /docker-main-app
elif [[ $ENVIRONMENT_MODE == "development" ]]
then
    go get github.com/githubnemo/CompileDaemon
    CompileDaemon --build="go build -o /docker-main-app ./cmd/main" --command=/docker-main-app
fi
