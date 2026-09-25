FROM ubuntu:26.04

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
RUN curl -fL https://download.swift.org/swift-6.4.0-release/ubuntu2604-aarch64/swift-6.4.0-RELEASE/swift-6.4.0-RELEASE-ubuntu26.04-aarch64.tar.gz -o swift.tar.gz
RUN tar xzf swift.tar.gz --directory / --strip-components=1 && rm swift.tar.gz

WORKDIR /app

# Build swift-protobuf from source
# Note: --recursive is required to fetch C++ protobuf submodules
RUN git clone --recursive https://github.com/apple/swift-protobuf.git && \
    cd swift-protobuf && \
    swift build -c release && \
    cp .build/release/protoc-gen-swift /usr/local/bin/ && \
    cd .. && \
    rm -rf swift-protobuf

# Build protoc-gen-grpc-swift from source
# Note(@adriano): grpc-swift 1.x does not support Swift 6.4, so use the
# grpc-swift-protobuf v2 code generator instead.
# Install protoc-gen-grpc-swift-2 as protoc-gen-grpc-swift so our existing
# --grpc-swift_out invocation continues to work.
RUN git clone --branch 2.4.0 https://github.com/grpc/grpc-swift-protobuf.git && \
    cd grpc-swift-protobuf && \
    swift build -c release --product protoc-gen-grpc-swift-2 && \
    cp .build/release/protoc-gen-grpc-swift-2 /usr/local/bin/protoc-gen-grpc-swift && \
    cd .. && \
    rm -rf grpc-swift-protobuf

# Need to download well-known protobuf types (e.g. timestamp, struct, ...)
RUN git clone https://github.com/protocolbuffers/protobuf.git && \
    mkdir -p /usr/include && \
    cp -r protobuf/src/* /usr/include/ && \
    rm -rf protobuf

VOLUME /app/gen

# NOTE(@adrianosela): Using shell form instead of exec form because we rely on
# the shell expanding the wildcard character '*' to find all .proto files in
# the given directory. This is discouraged and produces a warning. See for more
# info: https://docs.docker.com/reference/dockerfile/#shell-and-exec-form
ENTRYPOINT mkdir -p /app/gen/swift && protoc \
    -I/app/proto \
    -I/app/shared \
    --swift_out=/app/gen/swift \
    --grpc-swift_out=/app/gen/swift \
    /app/proto/*.proto
