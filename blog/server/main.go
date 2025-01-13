/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 13.01.2025
	Time        : 11:29
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.BlogServiceServer
}

var addr string = "0.0.0.0:50051"

var collection *mongo.Collection

func main() {
	// mongo
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog") // if doesn't exist, it will be created

	// This is a server
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer list.Close()

	log.Printf("server listening at: %s\n", addr)

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{}) // To implement BlogServiceServer interface -> implement in the blog.go file

	err = server.Serve(list)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
