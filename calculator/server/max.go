package main

import (
	pb "grpc-go/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalcService_MaxServer) error {
	log.Println("Max invoked.")

	var currMax int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error reading from client stream: %v\n", err)
		}

		if req.In > currMax {
			currMax = req.In
		}
		err = stream.Send(&pb.MaxResponse{Out: currMax})
		if err != nil {
			log.Fatalf("Error sending response to client: %v\n", err)
		}
	}
}
