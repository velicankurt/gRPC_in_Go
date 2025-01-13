/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 13.01.2025
	Time        : 11:50
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("ReadBlog function was invoked")
	blogId := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), blogId)
	if err != nil {
		log.Fatalf("failed to read blog: %v", err)
	}

	log.Printf("blog read: %v\n", res)
	return res
}
