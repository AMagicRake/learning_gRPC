package main

import (
	"context"
	pb "grpc-go/greet/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("Do greeat with deadline was invoked with deadline: %v\n", timeout)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Niel"}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded.")
				return
			} else {
				log.Fatalf("Unexpected grpc error: %v\n", err)
			}
		} else {
			log.Fatalln("Error from server: %v\n", err)
		}
	}

	log.Printf("Greet with deadline: %s\n", res.Result)
}
