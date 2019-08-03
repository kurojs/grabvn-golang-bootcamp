.SILENT:

generate:
	protoc -I proto/ --proto_path=$(HOME)/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis proto/user.proto --go_out=plugins=grpc:proto
	protoc -I proto/ --proto_path=$(HOME)/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:proto proto/user.proto

	echo "Generate done"

run-server:
	go run ./server/main.go

run-client:
	go run ./client/main.go