/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 15:28
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:50051"

func main() {
	// This is client
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c, 3, 10)
	doPrimes(c, 120)
	defer conn.Close()
}
