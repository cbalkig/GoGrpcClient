package language_id

import (
	grpcCaller "GoGrpcClient/grpc"
	languageId "GoGrpcClient/grpc/proto/clients/language_id/pb"
	"golang.org/x/net/context"
	"log"
)

func Call(text string) string {
	conn := grpcCaller.Init()

	c := languageId.NewLanguageIdServiceClient(conn)

	response, err := c.DetectFast(context.Background(), &languageId.LanguageIdRequest{Input: text})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.LangId)
	return response.LangId
}
