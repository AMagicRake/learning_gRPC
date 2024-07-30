package main

import (
	"context"
	pb "grpc-go/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Create blog was invoked")

	blog := &pb.Blog{
		AuthorId: "Niel",
		Title:    "My First Blog",
		Content:  "Content of the first blog.",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created %s\n", res.Id)
	return res.Id

}

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("Read blog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while requesting Blog: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}

func updateBlog(c pb.BlogServiceClient, toUpdate *pb.Blog) {
	log.Println("Update blog was invoked.")

	res, err := c.UpdateBlog(context.Background(), toUpdate)

	if err != nil {
		log.Printf("Error while updating blog: %v\n", err)
	}

	log.Printf("Blog was updated: %v\n", res)
}

func listBlog(c pb.BlogServiceClient) {
	log.Println("List blogs was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error while creating stream to list blogs: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Unknown error while reading stream: %v\n", err)
		}
		log.Printf("Blog: %v\n", msg)
	}
}
