package main

import (
	"context"
	pb "grpc-go/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("Greet Everone invoked.")

	reqs := []*pb.GreetRequest{
		{FirstName: "Niel"},
		{FirstName: "Megan"},
		{FirstName: "Gemma"},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating server stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error in reading stream: %v\n", err)
				break
			}
			log.Printf("Received %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
