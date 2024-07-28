package main

import (
	pb "grpc-go/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalcService_AvgServer) error {
	log.Println("Avg invoked")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//do average calc
			result := float64(sum) / float64(count)
			return stream.SendAndClose(&pb.AvgResponse{Output: result})
		}
		if err != nil {
			log.Fatalf("Error while receiving stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		sum += int(req.Input)
		count++
	}

}
