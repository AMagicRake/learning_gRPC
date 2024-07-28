package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalcServiceClient, n int32) {
	log.Printf("do Sqrt was invoked with: %d\n", n)

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We sent an invalid argument.")
				return
			}

		} else {
			log.Printf("Non gRPC error: %v\n", err)
		}
	}

	log.Printf("SquareRoot %f\n", res.Result)
}
