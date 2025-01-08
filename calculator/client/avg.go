/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 21:58
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient, avgRequests []*pb.AvgRequest) {
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while calling Avg RPC: %v", err)
	}

	for _, req := range avgRequests {
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending request to Avg: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from Avg: %v", err)
	}

	log.Printf("Avg Response: %v\n", res.Avg)
}
