build:
	@go build -o bin/main cmd/*.go
test:
	@go test -v ./...
run: build
	@REDIS_ADDR=localhost:6379 ./bin/main