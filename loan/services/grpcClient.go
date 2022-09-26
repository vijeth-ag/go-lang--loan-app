package services

import (
	"log"

	"google.golang.org/grpc"
)

var GrpcClient *grpc.ClientConn

func GrpcConn() *grpc.ClientConn {
	GrpcClient, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Println("err establishing connection to grpc on :9001")
	}
	return GrpcClient
}
