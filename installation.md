# installation

1. Protocol Buffer Compiler Installation: https://grpc.io/docs/protoc-installation/

2. install protoc-gen-go-grpc

```sh
$ export GO111MODULE=on  # Enable module mode
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

install protoc-gen-go without vpn:

```sh
$ go env -w GO111MODULE=on

$ go env -w GOPROXY="https://goproxy.cn,direct" 
# default $GOPATH/bin
$ go get -u github.com/golang/protobuf/protoc-gen-go

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

3. generate code

```sh
$ make genproto 
```

the output might be different under different versions of protoc or protoc-gen-go, we could use a fixed version in the dockerfile to keep the consistency, for example:

```dockerfile
ENV PROTOC_VERSION 3.6.1
ENV PROTOC_GEN_GO_VERSION v1.3.2
RUN apt-get update; \
    apt-get install -y unzip; \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip; \
    unzip -o protoc-$PROTOC_VERSION-linux-x86_64.zip -d /usr/local bin/protoc && \
    unzip -o protoc-$PROTOC_VERSION-linux-x86_64.zip -d /usr/local include/* && \
    rm -rf protoc-$PROTOC_VERSION-linux-x86_64.zip; \
    go get -u github.com/golang/protobuf/protoc-gen-go@$PROTOC_GEN_GO_VERSION; \
    cp "$(go env GOPATH)"/bin/protoc-gen-go /usr/bin;
```
