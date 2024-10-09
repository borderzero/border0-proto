ARG UBUNTU_VERSION=24.04
ARG ARCHITECTURE=aarch64
ARG GO_VERSION=v1.23
ARG SWIFT_VERSION=6.0.1
ARG PROTOC_GEN_SWIFT_VERSION=1.28.2
ARG PROTOC_GEN_SWIFT_GRPC_VERSION=1.24.1

FROM ubuntu:${UBUNTU_VERSION}

RUN apt-get update && apt-get install -y \
    curl \
    git \
    unzip \
    build-essential \
    clang \
    libssl-dev \
    protobuf-compiler

# install swift
RUN curl https://download.swift.org/swift-6.0.1-release/ubuntu2404-aarch64/swift-6.0.1-RELEASE/swift-6.0.1-RELEASE-ubuntu24.04-aarch64.tar.gz -o swift.tar.gz
RUN tar xzf swift.tar.gz --directory / --strip-components=1 && rm swift.tar.gz

WORKDIR /app

ARG PROTOC_GEN_GO_VERSION=v1.34.2
ARG PROTOC_GEN_GO_GRPC_VERSION=v1.5.1

# install go
RUN apt-get install -y golang-go
RUN go version
# install go protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOC_GEN_GO_VERSION}
# install go grpc
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOC_GEN_GO_GRPC_VERSION}

# RUN swift build --product protoc-gen-swift
# RUN swift build --product protoc-gen-grpc-swift

# # install swift protobuf
RUN git clone https://github.com/apple/swift-protobuf.git && \
    cd swift-protobuf && \
    swift build -c release && \
    cp .build/release/protoc-gen-swift /usr/local/bin/ && \
    cd .. && \
    rm -rf swift-protobuf
# install swift grpc
# RUN git clone https://github.com/grpc/grpc-swift.git && \
#     cd grpc-swift && \
#     swift build -c release # && \
#     # cp .build/release/protoc-gen-grpc-swift /usr/local/bin/ && \
#     # cd .. && \
#     # rm -rf grpc-swift

# need to download well-known types (e.g. timestamp, struct, ...)
RUN git clone https://github.com/protocolbuffers/protobuf.git && \
    mkdir -p /usr/include && \
    cp -r protobuf/src/* /usr/include/ && \
    rm -rf protobuf

VOLUME [ "/app/proto" ]

COPY generate.sh /generate.sh

ENTRYPOINT [ "/generate.sh" ]
