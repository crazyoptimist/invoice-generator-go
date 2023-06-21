APP_NAME=invoicegen
VERSION=1.0.0

run:
	go run cmd/$(APP_NAME)/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o build/linux/$(APP_NAME) cmd/$(APP_NAME)/main.go
	GOOS=windows GOARCH=amd64 go build -o build/windows/$(APP_NAME).exe cmd/$(APP_NAME)/main.go

clean:
	rm -rf build

.PHONY:
	build clean
