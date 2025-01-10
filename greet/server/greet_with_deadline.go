/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 10.01.2025
	Time        : 13:23
	Notes       :

*/

package main

import (
	"context"
	"errors"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadLine(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadLine function was invoked with %v", in)
	for i := 0; i < 3; i++ {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("The client has canceled the request")
			return nil, status.Error(codes.Canceled, "The client has canceled the request")
		}
		log.Printf("Sleeping for 1 second")
		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil
}
