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
	pb "github.com/velicankurt/gRPC_in_Go/blog/proto"
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
	c := pb.NewBlogServiceClient(conn)

	// CreateBlog
	blogID := createBlog(c)

	// ReadBlog
	readBlog(c, blogID) // valid id
	//readBlog(c, "invalid-id") // invalid id

	// UpdateBlog
	updateBlog(c, blogID)

	// DeleteBlog
	deleteBlog(c, blogID)

	// ListBlogs
	listBlogs(c)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(conn)
}
