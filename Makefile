SRC_PATH := ./src/cmd/main.go

.PHONE=run-dev

run-dev:
	go run $(SRC_PATH)