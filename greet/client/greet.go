package main

import (
	"context"
	pb "grpc-go/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {

	log.Println("doGreet service was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Niel"})

	if err != nil {
		log.Fatalf("Failed to call greet endpoint: %v\n", err)
	}

	log.Printf("Greeting: %s", res.Result)
}
