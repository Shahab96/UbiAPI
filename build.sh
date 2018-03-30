#! /bin/bash

# You may need to run chmod +x on this file before it can be executed.

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./prod/crud ./src/
docker-compose up --build -d
