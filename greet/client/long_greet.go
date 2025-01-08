/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 21:35
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient, greetRequests []*pb.GreetRequest) {
	// LongGreet function
	log.Println("LongGreet function was invoked")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet RPC: %v", err)
	}

	for _, req := range greetRequests {
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	log.Printf("LongGreet Response: %v\n", res.Result)
}
