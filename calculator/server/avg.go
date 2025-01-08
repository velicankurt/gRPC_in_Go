/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 21:50
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked with a streaming request")
	var sum int32
	var count int32
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Avg: float64(sum) / float64(count)})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}

		sum += recv.Number
		count++
	}
	return nil
}
