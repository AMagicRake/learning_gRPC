package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalcServiceClient) {
	log.Println("DoMax invoked.")

	reqs := []*pb.MaxRequest{
		{In: 8},
		{In: 15},
		{In: 8},
		{In: 3},
		{In: 34},
		{In: 67},
		{In: 3},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream to server: %v\n")
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
				log.Printf("Error while reading server stream: %v\n", err)
				break
			}
			log.Printf("Current max: %v\n", res.Out)
		}
		close(waitc)
	}()

	<-waitc
}
