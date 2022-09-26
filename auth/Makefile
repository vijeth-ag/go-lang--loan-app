maketest:
	echo "Hello World"

generate-proto:
	protoc --go_out=./proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./proto ./proto/auth.proto