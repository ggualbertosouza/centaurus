PLAYGROUND_PATH := ./src/cmd/playground/main.go

.PHONE=run-dev

run-dev:
	go run $(PLAYGROUND_PATH)

run-test:
	go test ./...