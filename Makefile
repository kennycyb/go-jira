#!/usr/bin/make -f
init:
	go get -u github.com/rogpeppe/gohack

generate:
	go generate ./...

test:
	go vet ./...
	go test ./...
