FROM golang:1.23-alpine3.20

RUN apk --no-cache update
RUN apk --no-cache upgrade
RUN apk --no-cache add protobuf git

WORKDIR /app

# Install protoc-gen-go, track versions in
# https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go?tab=versions
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2

# Install protoc-gen-go-grpc, track versions in
# https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc?tab=versions
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

# Need to download well-known protobuf types (e.g. timestamp, struct, ...)
RUN git clone https://github.com/protocolbuffers/protobuf.git && \
    mkdir -p /usr/include && \
    cp -r protobuf/src/* /usr/include/ && \
    rm -rf protobuf

VOLUME /app/proto

# NOTE(@adrianosela): Using shell form instead of exec form because we rely on
# the shell expanding the wildcard character '*' to find all .proto files in
# the given directory. This is discouraged and produces a warning. See for more
# info: https://docs.docker.com/reference/dockerfile/#shell-and-exec-form
ENTRYPOINT protoc \
    -I/app/proto \
    -I/app/shared \
    --go_out=paths=source_relative:/app/proto \
    --go-grpc_out=paths=source_relative:/app/proto \
    /app/proto/*.proto

