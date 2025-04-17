proto-google:
	curl https://raw.githubusercontent.com/googleapis/googleapis/974ad5bdfc9ba768db16b3eda2850aadd8c10a2c/google/api/annotations.proto --create-dirs -o ../google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/974ad5bdfc9ba768db16b3eda2850aadd8c10a2c/google/api/http.proto --create-dirs -o ../google/api/http.proto

proto-openapiv2:
	curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/annotations.proto \
		--create-dirs -o ../protoc-gen-openapiv2/options/annotations.proto
	curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/openapiv2.proto \
		--create-dirs -o ../protoc-gen-openapiv2/options/openapiv2.proto

build_proto: proto-google proto-openapiv2
	protoc -I . -I .. -I../protoc-gen-openapiv2/options \
		--go_out=../generated/clients --go_opt=paths=source_relative \
		--go-grpc_out=../generated/clients --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=../generated/clients --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=logtostderr=true:../generated/clients \
		clients.proto



