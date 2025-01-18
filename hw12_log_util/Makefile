env-prepare:
	cp -n .env.example .env

audit:
	golangci-lint run ./...

tidy:
	go mod tidy
	go fmt ./...
	golangci-lint run --fix ./...

run:
	go run main.go