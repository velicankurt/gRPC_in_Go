/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 22:16
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"io"
	"log"
	"sync"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient, greetRequests []*pb.GreetRequest) {
	// GreetEveryone function
	log.Println("GreetEveryone function was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while calling GreetEveryone RPC: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(2) // we have 2 go routines
	// we send a bunch of messages to the server (go routine)
	go func() {
		defer wg.Done()
		for _, req := range greetRequests {
			log.Printf("Sending request: %v\n", req)
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending request to GreetEveryone: %v", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		err = stream.CloseSend()
	}()

	// we receive a bunch of messages from the server (go routine)
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
	}()

	wg.Wait() // wait until all go routines are done
}
