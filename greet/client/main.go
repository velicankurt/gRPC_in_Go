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
	"google.golang.org/grpc/credentials"
	"log"
)

var addr string = "localhost:50051"

func main() {
	// This is a client
	tls := true // change to false if needed
	var options []grpc.DialOption
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("failed to load certificates: %v", err)
		}

		options = append(options, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(addr, options...) // grpc.Dial
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c, []*pb.GreetRequest{
	//	{FirstName: "Veli Can"},
	//	{FirstName: "Veli"},
	//	{FirstName: "Can"},
	//})
	//doGreetEveryone(c, []*pb.GreetRequest{
	//	{FirstName: "Veli"},
	//	{FirstName: "Can"},
	//	{FirstName: "Kurt"},
	//})
	//doGreetWithDeadLine(c, 5*time.Second)
	//doGreetWithDeadLine(c, 1*time.Second)
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)
}
