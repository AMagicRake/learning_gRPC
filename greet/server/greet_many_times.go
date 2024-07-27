package main

import (
	"fmt"
	pb "grpc-go/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes func invoked with: %v\n", in)

	for i := range 10 {
		res := fmt.Sprintf("Hellow %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
