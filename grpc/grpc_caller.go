package grpcCaller

import (
	zemberek_normalization "WebServiceDockerDemo/grpc/proto/clients/normalization"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Call() string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := zemberek_normalization.NewNormalizationServiceClient(conn)

	response, err := c.Normalize(context.Background(), &zemberek_normalization.NormalizationRequest{Input: "naaaaaber"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.NormalizedInput)
	return response.NormalizedInput
}