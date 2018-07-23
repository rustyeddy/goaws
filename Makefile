cmd = goa/goa

dirs = cmd

goa: *.go goa/*.go cmd/*.go
	make -C goa

build:
	go build

test:
	go test

run: 
	go run -v *.go

install:
	go install -v -o $(cmd)

.PHONY: build test run install
