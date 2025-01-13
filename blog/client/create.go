/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 13.01.2025
	Time        : 11:41
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("CreateBlog function was invoked")
	blog := &pb.Blog{
		AuthorId: "Veli Can",
		Title:    "My First Blog",
		Content:  "This is my first blog. I am so excited to write this blog.",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("failed to create blog: %v", err)
	}

	log.Printf("blog created with id: %s\n", res.Id)
	return res.Id
}
