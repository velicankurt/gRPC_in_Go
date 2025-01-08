/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 22:40
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"log"
	"sync"
	"time"
)

func doMax(c pb.CalculatorServiceClient, maxRequests []*pb.MaxRequest) {
	log.Println("doMax function was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error while calling Max RPC: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go func() { // sender
		defer wg.Done()
		for _, req := range maxRequests {
			log.Printf("Sending request: %v\n", req)
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending request to Max: %v", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("error while closing stream: %v", err)
		}
	}()

	go func() { // receiver
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				break
			}
			log.Printf("Received: %v\n", res.Maximum)
		}
	}()

	wg.Wait()
}
