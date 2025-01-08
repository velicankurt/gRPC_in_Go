/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 13:58
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	// This is a client
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.Dial
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	c := pb.NewGreetServiceClient(conn)
	//doGreet(c)
	doGreetManyTimes(c)
	defer conn.Close()
}
