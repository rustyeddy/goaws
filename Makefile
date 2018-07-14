cmd = goa

dirs = cmd

build:
		make -C goa build

test:
		go test

run: 
		go run -v *.go

install:
		go install -v -o $(cmd)

.PHONY: build test run install
