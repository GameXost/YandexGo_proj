proto-google:
	curl https://raw.githubusercontent.com/googleapis/googleapis/974ad5bdfc9ba768db16b3eda2850aadd8c10a2c/google/api/annotations.proto --create-dirs -o API/google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/974ad5bdfc9ba768db16b3eda2850aadd8c10a2c/google/api/http.proto --create-dirs -o API/google/api/http.proto

proto-openapiv2:
	curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/annotations.proto \
		--create-dirs -o API/protoc-gen-openapiv2/options/annotations.proto
	curl https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/master/protoc-gen-openapiv2/options/openapiv2.proto \
		--create-dirs -o API/protoc-gen-openapiv2/options/openapiv2.proto

generate_server:
	docker run --rm \
    -v ${CURDIR}:/local \
    openapitools/openapi-generator-cli generate \
    -i /local/API/generated/clients/clients.swagger.json \
    -g go-gin-server \
    -o /local/server \
    --ignore-file-override /local/API/generated/clients/.openapi-generator-ignore \
    --additional-properties=packageName=server \
    --git-user-id GameXost \
    --git-repo-id YandexGo_proj

generate_client_server:
	docker run --rm \
	-v ${CURDIR}:/local \
	openapitools/openapi-generator-cli generate \
	-i /local/API/generated/clients/clients.swagger.json \
	-g go \
	-o /local/go-client \
	--ignore-file-override /local/API/generated/clients/.openapi-generator-ignore \
	--additional-properties=packageName=go_client \
	--git-user-id GameXost \
	--git-repo-id YandexGo_proj

build_proto: proto-google proto-openapiv2 generate_server generate_client_server

	protoc -I API/protos -I API/ -I API/protoc-gen-openapiv2/options \
		--go_out=API/generated/clients --go_opt=paths=source_relative \
		--go-grpc_out=API/generated/clients --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=API/generated/clients --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=logtostderr=true:API/generated/clients \
		API/protos/clients.proto



docker_build:
	docker build \
		-t yandexgo_proj_users \
		-f Dockerfile \
		.

docker_run:
	docker run --rm -it \
		-p 9093:9093 -p 9094:9094 -p 9091:9091 \
		--name yandexgo_users_server \
		yandexgo_proj_users

docker_stop:
	docker stop yandexgo_users_server || true

run:
	go run ./cmd/main.go