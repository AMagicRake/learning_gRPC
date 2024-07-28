package main

import (
	"context"
	pb "grpc-go/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Do long greet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Niel"},
		{FirstName: "Megan"},
		{FirstName: "Gemma"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Long Greet: %s\n", res.Result)
}
