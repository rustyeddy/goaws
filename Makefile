cmd = goa

dirs = cmd

build:
		go build 
		make -C cmd build
		make -C goa build
		make -C goa run

test:
		go test

run: 
		go run -v *.go

install:
		go install -v -o $(cmd)

.PHONY: build test run install
