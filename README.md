# border0-proto

### Generating protobuf files

Simply run `make`

e.g.

```
12:58 $ make
docker build --build-arg PROTOC_GEN_GO_VERSION=v1.34.2 --build-arg PROTOC_GEN_GO_GRPC_VERSION=v1.5.1 -t border0-proto-builder .
[+] Building 0.4s (14/14) FINISHED                                                                                                                           docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                                                                         0.0s
 => => transferring dockerfile: 736B                                                                                                                                         0.0s
 => [internal] load metadata for docker.io/library/golang:1.22-alpine3.20                                                                                                    0.3s
 => [internal] load .dockerignore                                                                                                                                            0.0s
 => => transferring context: 2B                                                                                                                                              0.0s
 => [1/9] FROM docker.io/library/golang:1.22-alpine3.20@sha256:1a478681b671001b7f029f94b5016aed984a23ad99c707f6a0ab6563860ae2f3                                              0.0s
 => [internal] load build context                                                                                                                                            0.0s
 => => transferring context: 33B                                                                                                                                             0.0s
 => CACHED [2/9] RUN apk --no-cache update                                                                                                                                   0.0s
 => CACHED [3/9] RUN apk --no-cache upgrade                                                                                                                                  0.0s
 => CACHED [4/9] RUN apk --no-cache add protobuf git                                                                                                                         0.0s
 => CACHED [5/9] WORKDIR /app                                                                                                                                                0.0s
 => CACHED [6/9] RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2                                                                                         0.0s
 => CACHED [7/9] RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1                                                                                         0.0s
 => CACHED [8/9] RUN git clone https://github.com/protocolbuffers/protobuf.git &&     mkdir -p /usr/include &&     cp -r protobuf/src/* /usr/include/ &&     rm -rf protobu  0.0s
 => CACHED [9/9] COPY generate.sh /generate.sh                                                                                                                               0.0s
 => exporting to image                                                                                                                                                       0.0s
 => => exporting layers                                                                                                                                                      0.0s
 => => writing image sha256:ccdea7188d20715ac639bed2c7fa8fb117d8cf27270611670b0f7b4ba8e84cc8                                                                                 0.0s
 => => naming to docker.io/library/border0-proto-builder                                                                                                                     0.0s
docker run -v /Users/adriano/go/src/github.com/borderzero/border0-proto/connector:/app/proto border0-proto-builder; 	docker run -v /Users/adriano/go/src/github.com/borderzero/border0-proto/peer:/app/proto border0-proto-builder;
```
