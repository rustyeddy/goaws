cmd = goa

dirs = cmd srv

goa: main.go ../*.go ../cmd/*.go
	go build -o goa

build:
	go build -o $(cmd) 

inst: goa
	go run main.go instances

vols: goa
	go run main.go volumes

test:
	go test 

run: 
	go run *.go

install:
	go install -o $(cmd)

.PHONY: build test run install
