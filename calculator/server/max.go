/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 22:35
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	var maximum int32 = 0
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}

		if recv.Number > maximum {
			maximum = recv.Number
			log.Println("New maximum found: ", maximum)
			err = stream.Send(&pb.MaxResponse{Maximum: maximum})
			if err != nil {
				log.Fatalf("error while sending data to client: %v", err)
			}
		}
	}
}
