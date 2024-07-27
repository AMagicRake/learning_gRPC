package main

import (
	pb "grpc-go/calculator/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50052"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	// doAdd(c)
	doPrimes(c)
}
