build:
	go build server.go

build-linux:
	env GOOS=linux go build -o server-linux server.go 
