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
	"google.golang.org/grpc/credentials"
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

	var options []grpc.ServerOption
	tls := true // chenge to false if needed
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to load certificates: %v", err)
		}
		options = append(options, grpc.Creds(creds))
	}

	server := grpc.NewServer(options...)             // ssl implementation
	pb.RegisterGreetServiceServer(server, &Server{}) // To implement GreetServiceServer interface -> implement in the greet.go file
	err = server.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
