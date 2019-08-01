.SILENT:

generate:
	protoc -I proto/ proto/feedback.proto --go_out=plugins=grpc:proto
	echo "Generate done"

run-server:
	go run ./server/main.go

run-client:
	go run ./client/main.go