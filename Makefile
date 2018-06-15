cmd = goaws

dirs = cmd srv

build:
		go build -o $(cmd) -v .

test:
		go test -v 

run: 
		go run -v main.go
