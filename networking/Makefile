.PHONY: build run clean proto-gen

GOOGLEAPIS := $(shell go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis

proto-gen:
	protoc --go_out=gen/pb --go_opt=paths=source_relative \
		--go-grpc_out=gen/pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=gen/pb --grpc-gateway_opt=paths=source_relative \
		-I api/proto \
		-I $(shell brew --prefix)/include \
		-I $(GOOGLEAPIS) \
		api/proto/user.proto

build:
	go build -o bin/server ./cmd

run: build
	./bin/server

clean:
	rm -rf bin/
