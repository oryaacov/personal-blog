build:
	GOOS=linux GOARCH=amd64 go build -o bin/server ./cmd/server/server.go

build-all:
	GOOS=windows go build -o ./bin/server.exe ./cmd/server/server.go
	GOOS=linux GOARCH=amd64 go build -o bin/server ./cmd/server/server.go