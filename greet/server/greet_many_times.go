/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 16:05
	Notes       :

*/

package main

import (
	"fmt"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s for the %d. time", in.FirstName, i)
		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
