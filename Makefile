ver?=v1

dev:
	go run main.go routes.go

grpc:
	$(eval PROTO_PATH = ./pkg/server/$(ver)/grpc/models)
	mkdir -p "./pkg/server/api_pb"
	protoc \
		-I. \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--grpc-gateway_out=logtostderr=true,allow_delete_body=true,paths=source_relative:. \
		--swagger_out=allow_merge=true,merge_file_name=api:./pkg/server \
		--go_out=paths=source_relative:. /*.proto

