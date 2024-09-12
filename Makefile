build:
	@go build -o bin/main *.go
test:
	@go test -v ./...
run: build
	@./bin/main