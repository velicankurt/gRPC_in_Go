/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 8.01.2025
	Time        : 22:59
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) {
	log.Printf("Starting to do a Unary RPC for Sqrt")
	req := &pb.SqrtRequest{Number: number}
	res, err := c.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %v", e.Message())
			log.Printf("Error code from server: %v", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Fatalf("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("A non gRPC error: %v", err)
		}
	}
	log.Printf("Response from Sqrt: %v", res.Result)
}
