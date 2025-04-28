# Makefile

APP_NAME=server-project

run:
	go run .

build:
	go build -o $(APP_NAME)

clean:
	rm -f $(APP_NAME)

fmt:
	go fmt ./...

