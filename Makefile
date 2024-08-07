NAME:=$(shell basename `git rev-parse --show-toplevel`)

# track versions in https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go?tab=versions
PROTOC_GEN_GO_VERSION:=v1.34.2

# track versions in https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc?tab=versions
PROTOC_GEN_GO_GRPC_VERSION:=v1.5.1

component ?= connector

all: docker-run

docker-run: docker-build
	docker run -v $(PWD)/$(component):/app/proto -e PROTO_COMPONENT=$(component) $(NAME)-builder

docker-build:
	docker build --build-arg PROTOC_GEN_GO_VERSION=$(PROTOC_GEN_GO_VERSION) --build-arg PROTOC_GEN_GO_GRPC_VERSION=$(PROTOC_GEN_GO_GRPC_VERSION) -t $(NAME)-builder .
