#!/bin/sh

protoc -I/app/proto --proto_path=/app/protos --go_out=paths=source_relative:/app/proto --go-grpc_out=paths=source_relative:/app/proto /app/proto/*.proto
