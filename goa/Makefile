cmd = goaws

dirs = cmd srv

build:
		go build -o $(cmd) 

test:
		go test 

run: 
		go run *.go

install:
		go install -o $(cmd)

.PHONY: build test run install
