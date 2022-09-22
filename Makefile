ver?=v1

dev:
	go run main.go routes.go

generate: generate_grpc generate_client
generate_grpc:
	$(eval PROTO_PATH = ./pkg/server/$(ver)/grpc/models)
	protoc \
		-I. \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true,allow_delete_body=true,paths=source_relative:. \
		--swagger_out=allow_merge=true,merge_file_name=api:$(PROTO_PATH) \
		--go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $(PROTO_PATH)/*.proto

generate_client:
	swagger generate client -f ./pkg/server/v1/grpc/models/api.swagger.json -A year-wheel -t ./pkg/server/v1/grpc/clients

startdb:
	mongod --dbpath ./data/dev/mongo/
