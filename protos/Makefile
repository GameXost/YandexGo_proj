build_proto:
	protoc --go_out=./Clients/clients_proto --go_opt=paths=source_relative \
	 --go-grpc_out=./Clients/clients_proto --go-grpc_opt=paths=source_relative \
	  clients.proto
	protoc --go_out=./Drivers/drivers_proto --go_opt=paths=source_relative \
	 --go-grpc_out=./Drivers/drivers_proto --go-grpc_opt=paths=source_relative \
	  clients.proto