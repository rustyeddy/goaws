cmd = goaws

dirs = cmd srv

build:
		go build -o $(cmd) -v .

test:
		go test -v 

run: 
		go run -v *.go

install:
		go install -v -o $(cmd)

.PHONY: build test run install
