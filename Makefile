dirs = cmd srv

build:
		go build -v .

test:
		go test -v 

run: 
		go run -v main.go
