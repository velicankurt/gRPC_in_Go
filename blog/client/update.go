/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 13.01.2025
	Time        : 13:12
	Notes       :

*/

package main

import (
	"context"
	pb "github.com/velicankurt/gRPC_in_Go/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("UpdateBlog function was invoked")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Veli Can Kurt",
		Title:    "My First New Blog",
		Content:  "This is my first new blog. I am so excited to write this blog.",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("failed to update blog: %v", err)
	}

	log.Println("blog updated")
}
