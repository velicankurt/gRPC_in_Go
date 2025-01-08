/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 16:18
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	var k int32 = 2
	for in.Number > 1 {
		if in.Number%k == 0 {
			stream.Send(&pb.PrimesResponse{Prime: k})
			in.Number /= k
		} else {
			k++
		}
	}
	return nil
}
