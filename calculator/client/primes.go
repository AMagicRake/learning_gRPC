package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalcServiceClient) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimesRequest{Input: 120}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error when executing primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error when receiving message: %v\n", err)
		}

		log.Printf("Prime factor: %d\n", msg.Output)
	}
}
