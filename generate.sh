#!/bin/sh

protoc -I/app/proto --proto_path=/app/shared \
#    --go_out=paths=source_relative:/app/proto --go-grpc_out=paths=source_relative:/app/proto \        # go
    --swift_out=paths=source_relative:/app/proto --grpc-swift_out=paths=source_relative:/app/proto \  # swift
    /app/proto/*.proto
