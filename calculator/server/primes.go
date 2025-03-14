package main

import pb "grpc-go/calculator/proto"

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalcService_PrimesServer) error {
	var k int32 = 2
	primeNumber := in.Input

	for primeNumber > 1 {

		if primeNumber%k == 0 {
			stream.Send(&pb.PrimesResponse{Output: k})
			primeNumber /= k
		} else {
			k++
		}

	}
	return nil
}
