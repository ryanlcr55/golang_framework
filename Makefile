# injection
wire:
	@ go get github.com/google/wire/cmd/wire@v0.5.0
	@ cd internal/app; go run github.com/google/wire/cmd/wire gen

proto:
	@ protoc api/protobuf/*.proto --go_out=./internal --go-grpc_out=./internal


openapi_gen:
	@./scripts/openapi_gen.sh