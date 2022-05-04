.PHONY: build

unit-tests:
	go fmt ./...
	go test -v ./...

run:
	go run ./cmd/go-bricks