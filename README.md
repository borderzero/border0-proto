# border0-proto

### Generating protobuf files

Simply run `make component=${COMPONENT}`

e.g. for `COMPONENT=connector``

```
10:29 $ make component=connector
docker build --build-arg PROTOC_GEN_GO_VERSION=v1.31.0 --build-arg PROTOC_GEN_GO_GRPC_VERSION=v1.3.0 -t border0-proto-builder .
[+] Building 1.0s (14/14) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                      0.0s
 => => transferring dockerfile: 37B                                                                                                                                       0.0s
 => [internal] load .dockerignore                                                                                                                                         0.0s
 => => transferring context: 2B                                                                                                                                           0.0s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine3.18                                                                                                 0.9s
 => [1/9] FROM docker.io/library/golang:1.20-alpine3.18@sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c                                           0.0s
 => => resolve docker.io/library/golang:1.20-alpine3.18@sha256:fd9d9d7194ec40a9a6ae89fcaef3e47c47de7746dd5848ab5343695dbbd09f8c                                           0.0s
 => [internal] load build context                                                                                                                                         0.0s
 => => transferring context: 33B                                                                                                                                          0.0s
 => CACHED [2/9] RUN apk --no-cache update                                                                                                                                0.0s
 => CACHED [3/9] RUN apk --no-cache upgrade                                                                                                                               0.0s
 => CACHED [4/9] RUN apk --no-cache add protobuf git                                                                                                                      0.0s
 => CACHED [5/9] WORKDIR /app                                                                                                                                             0.0s
 => CACHED [6/9] RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0                                                                                      0.0s
 => CACHED [7/9] RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0                                                                                      0.0s
 => CACHED [8/9] RUN git clone https://github.com/protocolbuffers/protobuf.git &&     mkdir -p /usr/include &&     cp -R protobuf/src/* /usr/include/ &&     rm -rf prot  0.0s
 => CACHED [9/9] COPY generate.sh /generate.sh                                                                                                                            0.0s
 => exporting to image                                                                                                                                                    0.0s
 => => exporting layers                                                                                                                                                   0.0s
 => => writing image sha256:3b569ee4c24e98107e5a06c091532e1223022ec9ba9d4ed29403dfc0f5675e0c                                                                              0.0s
 => => naming to docker.io/library/border0-proto-builder                                                                                                                  0.0s
docker run -v /Users/adriano/go/src/github.com/borderzero/border0-proto/connector:/app/proto border0-proto-builder

```
