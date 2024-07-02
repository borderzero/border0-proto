FROM golang:1.22-alpine3.19

RUN apk --no-cache update
RUN apk --no-cache upgrade
RUN apk --no-cache add protobuf git

WORKDIR /app

ARG PROTOC_GEN_GO_VERSION=latest
ARG PROTOC_GEN_GO_GRPC_VERSION=latest

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VERSION}

# need to download well-known types (e.g. timestamp, struct, ...)
RUN git clone https://github.com/protocolbuffers/protobuf.git && \
    mkdir -p /usr/include && \
    cp -r protobuf/src/* /usr/include/ && \
    rm -rf protobuf

VOLUME [ "/app/proto" ]

COPY generate.sh /generate.sh

ENTRYPOINT [ "/generate.sh" ]
