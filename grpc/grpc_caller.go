package grpcCaller

import (
	"log"

	"google.golang.org/grpc"
)

func Init() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:6789", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn
}
