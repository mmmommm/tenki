.PHONY: test
test:
		go test ./...
lint:
		gofmt -w ./...
build:
		cp ./dev.env .env