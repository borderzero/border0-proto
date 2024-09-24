NAME:=$(shell basename `git rev-parse --show-toplevel`)

# track versions in https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go?tab=versions
PROTOC_GEN_GO_VERSION:=v1.34.2

# track versions in https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc?tab=versions
PROTOC_GEN_GO_GRPC_VERSION:=v1.5.1

# add any components here after their directory has been created with .proto files
COMPONENTS:=connector device 

all: docker-run

docker-run: docker-build
	$(foreach component, $(COMPONENTS),	docker run -v $(PWD):/app/protos -v $(PWD)/$(component):/app/proto $(NAME)-builder;)

docker-build:
	docker build --build-arg PROTOC_GEN_GO_VERSION=$(PROTOC_GEN_GO_VERSION) --build-arg PROTOC_GEN_GO_GRPC_VERSION=$(PROTOC_GEN_GO_GRPC_VERSION) -t $(NAME)-builder .
