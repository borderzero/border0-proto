#!/bin/sh

if [ -z "${PROTO_COMPONENT}" ]; then
    echo "ERROR: env PROTO_COMPONENT was not defined"
    exit 1
fi

protoc -I/app/proto --go_out=paths=source_relative:/app/proto --go-grpc_out=paths=source_relative:/app/proto /app/proto/${PROTO_COMPONENT}.proto
