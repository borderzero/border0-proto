NAME:=$(shell basename `git rev-parse --show-toplevel`)
PROTOC_GEN_GO_VERSION:=v1.31.0
PROTOC_GEN_GO_GRPC_VERSION:=v1.3.0

component ?= connector

all: docker-run

docker-run: docker-build
	docker run -v $(PWD)/$(component):/app/proto -e PROTO_COMPONENT=$(component) $(NAME)-builder

docker-build:
	docker build --build-arg PROTOC_GEN_GO_VERSION=$(PROTOC_GEN_GO_VERSION) --build-arg PROTOC_GEN_GO_GRPC_VERSION=$(PROTOC_GEN_GO_GRPC_VERSION) -t $(NAME)-builder .
