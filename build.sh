#!/bin/bash
GOOS=linux GOARCH=386 go build ./tiny-goweb.go
docker build -t kubeweb .
docker tag kubeweb drlee001/kubeweb
docker push drlee001/kubeweb
