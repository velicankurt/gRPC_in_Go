/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 22:56
	Notes       :

*/

package main

import (
	"context"
	"fmt"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("Sqrt function was invoked with a request")

	number := in.Number
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", number))
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
