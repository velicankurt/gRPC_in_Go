/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 16:18
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient, number int32) {
	// Primes function
	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{Number: number})
	if err != nil {
		log.Fatalf("error while calling Primes RPC: %v", err)
		return
	}

	// Print the result
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		println(response.Prime)
	}
}
