/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 21:27
	Notes       :

*/

package main

import (
	"fmt"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked with a streaming request")
	resp := ""
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: resp})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}

		resp += fmt.Sprintf("Hello %s!\n", recv.FirstName)
	}
}
