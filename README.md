# Now count the time

## Installation
```sh
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u google.golang.org/grpc
```

## Generate gRPC client and server
```now.proto
syntax = "proto3";
package now;

option go_package = "pb/now";

import "google/protobuf/timestamp.proto";
import "google/protobuf/empty.proto";

service Now {
    rpc Tick(google.protobuf.Empty) returns (stream google.protobuf.Timestamp) {}
}
```

check protoc-gen-go and protoc-gen-go-grpc version.
```sh
$ ~/go/bin/protoc-gen-go --version
protoc-gen-go v1.26.0
$ ~/go/bin/protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.0.1
```

generate 2 go files.
```sh
$ protoc --go_out=. --go-grpc_out=. proto/*.proto
# generated {go_package_path}/{proto_file_name}_grpc.pb.go
# generated {go_package_path}/{proto_file_name}.pb.go
```

## Run
```sh
$ go run server/server.go
$ go run client/client.go
```

#### Reference
https://blog.golang.org/protobuf-apiv2

#### what i don't understand
if i exec command below when cannot find packages, i can build and run.
```sh
$ go mod tidy
$ go mod vendor
```
