/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 15:26
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	// This is a server
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer list.Close()

	log.Printf("server listening at: %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &Server{}) // To implement CalculatorServiceServer interface -> implement in the calculator.go file
	err = server.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
