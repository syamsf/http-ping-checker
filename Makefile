build-linux:
	GOOS=linux GOARCH=amd64 go build -o web-checker

build:
	go build -o web-checker