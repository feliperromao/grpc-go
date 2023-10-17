generate:
	protoc --go_out=. --go-grpc_out=. proto/course_category.proto

run:
	go run cmd/grpcServer/maind.go 