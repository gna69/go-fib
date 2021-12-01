# go-fib

GRPC and HTTP REST server returns fibonacci sequence from X index to Y.

For compile protobuf use this command: `protoc --go_out=../../internal/api/grpc --go_opt=paths=source_relative --go-grpc_out=../../internal/api/grpc --go-grpc_opt=paths=source_relative fib.proto`;