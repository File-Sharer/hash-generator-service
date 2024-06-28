proto:
	protoc --go_out=. --go_opt=module=github.com/File-Sharer/hash-generator-service --go-grpc_out=. --go-grpc_opt=module=github.com/File-Sharer/hash-generator-service ./internal/proto/*.proto

dev:
	go run cmd/app/main.go

build:
	go build -o ./cmd/app/app ./cmd/app/main.go

run:
	./cmd/app/app
