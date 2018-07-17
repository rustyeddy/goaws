cmd = goa

dirs = cmd

build:
	go build

test:
	go test

run: 
	go run -v *.go

install:
	go install -v -o $(cmd)

.PHONY: build test run install
