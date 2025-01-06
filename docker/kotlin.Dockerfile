FROM ubuntu:24.04

# Set the default shell to bash
SHELL ["/bin/bash", "-c"]

# Update and install necessary packages
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    git \
    zip \
    openjdk-17-jdk \
    protobuf-compiler

WORKDIR /app

# Install protoc-gen-grpc-kotlin, track versions in
# https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-kotlin/
RUN curl -s https://repo1.maven.org/maven2/io/grpc/protoc-gen-grpc-kotlin/1.4.1/protoc-gen-grpc-kotlin-1.4.1-jdk8.jar -o /usr/local/bin/protoc-gen-grpc-kotlin.jar

# Need to create an executable wrapper for the Jar for use with protoc
RUN echo '#!/bin/bash' > /usr/local/bin/protoc-gen-grpc-kotlin.sh && \
    echo 'java -jar /usr/local/bin/protoc-gen-grpc-kotlin.jar "$@"' >> /usr/local/bin/protoc-gen-grpc-kotlin.sh && \
    chmod +x /usr/local/bin/protoc-gen-grpc-kotlin.sh

# Need to download well-known protobuf types (e.g., timestamp, struct, etc.)
RUN git clone https://github.com/protocolbuffers/protobuf.git && \
    mkdir -p /usr/include && \
    cp -r protobuf/src/* /usr/include/ && \
    rm -rf protobuf

VOLUME /app/gen

# NOTE(@adrianosela): Using shell form instead of exec form because we rely on
# the shell expanding the wildcard character '*' to find all .proto files in
# the given directory. This is discouraged and produces a warning. See for more
# info: https://docs.docker.com/reference/dockerfile/#shell-and-exec-form
ENTRYPOINT mkdir -p /app/gen/kotlin && protoc \
    -I/app/proto \
    -I/app/shared \
    --plugin=protoc-gen-kotlin_grpc=/usr/local/bin/protoc-gen-grpc-kotlin.sh \
    --kotlin_out=/app/gen/kotlin \
    --kotlin_grpc_out=/app/gen/kotlin \
    /app/proto/*.proto
