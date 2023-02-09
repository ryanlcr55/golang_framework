# injection
wire:
	@ cd internal/app; go run github.com/google/wire/cmd/wire gen

proto:
	@ protoc api/protobuf/*.proto --go_out=./internal --go-grpc_out=./internal