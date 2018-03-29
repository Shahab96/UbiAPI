#! /bin/bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./prod/crud ./src/
docker-compose up --build -d