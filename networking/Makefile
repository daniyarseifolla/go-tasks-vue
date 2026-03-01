.PHONY: build run clean proto-gen

proto-gen:
	protoc --go_out=pkg/pb --go_opt=paths=source_relative \
		--go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
		-I api/proto \
		-I $(shell brew --prefix)/include \
		api/proto/user.proto

build:
	go build -o bin/server .

run: build
	./bin/server

clean:
	rm -rf bin/
