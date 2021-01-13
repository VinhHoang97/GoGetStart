package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-exaple/chat"
	"log"
	"time"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.SayHello(ctx, &chat.MessageRequest{Code: "QN"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)
}
