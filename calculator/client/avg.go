package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"log"
)

func doAverage(c pb.CalcServiceClient) {
	log.Println("DoAverage invoked")

	reqs := []*pb.AvgRequest{
		{Input: 5},
		{Input: 8},
		{Input: 12},
		{Input: 32},
		{Input: 76},
		{Input: 50},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Printf("Error opening stream: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error receiving response: %v\n", err)
	}

	log.Printf("Average result: %d\n", res.Output)

}
