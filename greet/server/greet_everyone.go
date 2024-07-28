package main

import (
	pb "grpc-go/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Greet Everone invoked.")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error reading from stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
