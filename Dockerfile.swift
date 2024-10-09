FROM ubuntu:24.04

RUN apt-get update && apt-get install -y \
    curl \
    git \
    unzip \
    build-essential \
    clang \
    libssl-dev \
    protobuf-compiler

# Install swift, track versions in https://www.swift.org/download/
# Note(@adriano): I have an M1 MacBook Pro (arm), so I need the aarch64 version. If you are on an intel machine, use the x86_64 version.
RUN curl https://download.swift.org/swift-6.0.1-release/ubuntu2404-aarch64/swift-6.0.1-RELEASE/swift-6.0.1-RELEASE-ubuntu24.04-aarch64.tar.gz -o swift.tar.gz
RUN tar xzf swift.tar.gz --directory / --strip-components=1 && rm swift.tar.gz

WORKDIR /app

# Build swift-protobuf from source
RUN git clone https://github.com/apple/swift-protobuf.git && \
    cd swift-protobuf && \
    swift build -c release && \
    cp .build/release/protoc-gen-swift /usr/local/bin/ && \
    cd .. && \
    rm -rf swift-protobuf

# Build protoc-gen-grpc-swift from source
# Note(@adriano): Use the release/1.x branch... main is broken
# for Ubuntu. See https://github.com/grpc/grpc-swift/issues/2092.
RUN git clone -b release/1.x https://github.com/grpc/grpc-swift.git && \
    cd grpc-swift && \
    swift build -c release && \
    cp .build/release/protoc-gen-grpc-swift /usr/local/bin/ && \
    cd .. && \
    rm -rf grpc-swift

# Need to download well-known protobuf types (e.g. timestamp, struct, ...)
RUN git clone https://github.com/protocolbuffers/protobuf.git && \
    mkdir -p /usr/include && \
    cp -r protobuf/src/* /usr/include/ && \
    rm -rf protobuf

VOLUME /app/proto

ENTRYPOINT protoc \
    -I/app/proto \
    -I/app/shared \
    --swift_out=/app/proto \
    --grpc-swift_out=/app/proto \
    /app/proto/*.proto

