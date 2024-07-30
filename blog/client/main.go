package main

import (
	pb "grpc-go/blog/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50053"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	read := readBlog(c, id)
	readBlog(c, "blah")
	read.Content = "This is the updated content."
	updateBlog(c, read)
	readBlog(c, id)

}
