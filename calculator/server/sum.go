/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 15:31
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
