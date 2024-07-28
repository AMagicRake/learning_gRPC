package main

import (
	"log"
	"time"

	pb "grpc-go/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	// doGreet(c)
	// doGreetMany(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
	doGreetWithDeadline(c, 2*time.Second)
}
