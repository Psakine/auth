LOCAL_BIN:=$(CURDIR)/bin

build-mac:
	go build -o ${LOCAL_BIN}/ ${CURDIR}/cmd/grpc_server/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${LOCAL_BIN}/service_auth_linux ${CURDIR}/cmd/grpc_server/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) ${LOCAL_BIN}/golangci-lint run ./... --config .golangci.pipeline.yaml

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/psakine/auth-server:v0.0.1 .
	docker login -u token -p CRgAAAAASS16bb7CULZ5iZ4Yh8O1NwXIgzK4xJ6L cr.selcloud.ru/psakine
	docker push cr.selcloud.ru/psakine/auth-server:v0.0.1