package main

import (
	"context"
	pb "grpc-go/greet/proto"
	"io"
	"log"
)

func doGreetMany(c pb.GreetServiceClient) {
	log.Println("Do greet many was invoked.")

	req := &pb.GreetRequest{
		FirstName: "Niel",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Errror while reading stream %v\n", err)
		}
		log.Printf("Greet many times %s\n", msg.Result)
	}
}
