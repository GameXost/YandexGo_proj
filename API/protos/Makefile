build_proto:
	protoc -I . -I ../googleapis \
		--go_out=../generated/clients --go_opt=paths=source_relative \
		--go-grpc_out=../generated/clients --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../generated/clients --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=../generated/clients \
		clients.proto

	protoc -I . -I ../googleapis \
		--go_out=../generated/drivers --go_opt=paths=source_relative \
		--go-grpc_out=../generated/drivers --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../generated/drivers --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=../generated/drivers \
		drivers.proto
