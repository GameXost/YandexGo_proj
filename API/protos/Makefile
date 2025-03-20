build_proto:
	protoc --go_out=../generated/clients --go_opt=paths=source_relative \
	 --go-grpc_out=../generated/clients --go-grpc_opt=paths=source_relative \
	  clients.proto
	protoc --go_out=../generated/drivers --go_opt=paths=source_relative \
	 --go-grpc_out=../generated --go-grpc_opt=paths=source_relative \
	  drivers.proto