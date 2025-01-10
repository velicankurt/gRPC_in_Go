/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 10.01.2025
	Time        : 13:27
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadLine(c pb.GreetServiceClient, timeOut time.Duration) {
	log.Printf("doGreetWithDeadLine function was invoked with a deadline of %v\n", timeOut)

	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Veli Can"}
	res, err := c.GreetWithDeadLine(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Error message from server: %v", e.Message())
			}
		} else {
			log.Fatalf("A non gRPC error: %v", err)
		}
	}

	log.Printf("Response from GreetWithDeadLine: %v", res.Result)
}
