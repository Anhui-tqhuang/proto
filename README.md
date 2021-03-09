# proto

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
$ protoc -I. --go_out=plugins=grpc:./pd/auth ./auth.proto

$ protoc -I. --go_out=plugins=grpc:./pd/fight ./fight.proto
```
