package main

import (
	"context"
	"fmt"
	pb "grpc-go/calculator/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt invoked with: %v\n", in)

	number := in.Number
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Receivied a negative number: %d", number))
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
