/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 15:13
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
