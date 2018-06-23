cmd = goaws

dirs = cmd srv

build:
		go build -o $(cmd) .

test:
		go test

run: 
		go run -v *.go

install:
		go install -v -o $(cmd)

.PHONY: build test run install
