#!/bin/bash

docker-compose --file localenv/docker-compose.yml up -d --build
sleep 5
docker-compose --file localobservability/docker-compose.yml up -d --build 
sleep 10
go run ../bootstraper/bootstraper.go
