/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 7.01.2025
	Time        : 13:52
	Notes       :

*/

package main

import (
	pb "github.com/velicankurt/gRPC_in_Go/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer list.Close()

	log.Printf("server listening at: %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterGreetServiceServer(server, &Server{}) // To implement GreetServiceServer interface -> implement in the greet.go file
	err = server.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
