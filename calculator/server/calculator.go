package main

import (
	"context"
	pb "grpc-go/calculator/proto"
)

func (s Server) Add(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{Result: in.ValueX + in.ValueY}, nil
}
