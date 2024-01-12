PROTOC=protoc
PROTOC_FLAGS=--proto_path=./proto -I proto/googleapis
GENERATE_DIR=gen/go
MODULE_PATH=grpc-streaming-sample/grpc-server/gen/go

generate:
	$(PROTOC) $(PROTOC_FLAGS) \
		--go_out=$(GENERATE_DIR) \
		--go_opt=module=$(MODULE_PATH) \
		--go-grpc_out=$(GENERATE_DIR) \
		--go-grpc_opt=module=$(MODULE_PATH) \
		--grpc-gateway_out=$(GENERATE_DIR) \
		--grpc-gateway_opt=module=$(MODULE_PATH) \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=generate_unbound_methods=true \
		--openapiv2_out=gen/openapiv2 \
		--openapiv2_opt=allow_merge=true \
		proto/hello/*.proto

run_grpc_server:
	go run main.go
run_grpc_gateway:
	go run reverse_proxy/main.go