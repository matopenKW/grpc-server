export GOBIN=$(shell pwd)/bin

init:
	go get github.com/grpc-ecosystem/grpc-gateway@v1.16.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: protoc
protoc: ## gRPCのstubコードの生成
	protoc \
	-I=./proto \
	-I=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
	--go_out=./pkg/go \
	--go-grpc_out=require_unimplemented_servers=false:./pkg/go \
	proto/*/*.proto

	gofmt -s -w proto/
	gofmt -s -w pkg/go/
	goimports -w -local "github.com/grpc-server" pkg/go/