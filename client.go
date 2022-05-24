package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tutorialedge/go-grpc-tutorial/chat"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	
	response, err := c.SayTestName(context.Background(), &chat.Message{Body: "ut_registration.py"})
	if err != nil {
		log.Fatalf("Error when calling SayTestName: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
