/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : VCK
	Name        : Veli Can Kurt
	Date        : 13.01.2025
	Time        : 13:19
	Notes       :

*/

package main

import (
	"context"
	"fmt"
	pb "github.com/velicankurt/gRPC_in_Go/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs function was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("unknown internal error: %v", err))
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
		if err != nil {

		}
	}(cur, context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err = cur.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("error while decoding data: %v", err))
		}

		err = stream.Send(documentToBlog(data))
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("error while sending data to client: %v", err))
		}
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("unknown internal error: %v", err))
	}

	return nil
}
