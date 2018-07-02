cmd = goa

dirs = cmd srv

goa: main.go
	go build -o $@ $<

inst: goa
	go run main.go instances

vols: goa
	go run main.go volumes

build:
		go build -o $(cmd) 

test:
		go test 

run: 
		go run *.go

install:
		go install -o $(cmd)

.PHONY: build test run install
