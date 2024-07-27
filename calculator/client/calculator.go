package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"log"
)

func doAdd(c pb.CalcServiceClient) {

	res, err := c.Add(context.Background(), &pb.CalcRequest{ValueX: 5, ValueY: 7})
	if err != nil {
		log.Fatalf("Failed to call service: %v\n", err)
	}

	log.Printf("Result of addition: %v\n", res.Result)

}
