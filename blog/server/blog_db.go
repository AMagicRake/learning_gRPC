package main

import (
	"context"
	"fmt"
	pb "grpc-go/blog/proto"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateBlog(ctx context.Context, inData *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create Blog was invoked with: %v\n", inData)

	data := BlogItem{
		AuthorID: inData.AuthorId,
		Title:    inData.Title,
		Content:  inData.Content,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}

func (s *Server) ReadBlog(ctx context.Context, inData *pb.BlogId) (*pb.Blog, error) {
	log.Printf("Read Blog was invoked with: %v\n", inData)

	oid, err := primitive.ObjectIDFromHex(inData.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog item",
		)
	}
	return documentToBlog(data), nil
}

func (s *Server) UpdateBlog(ctx context.Context, inData *pb.Blog) (*empty.Empty, error) {
	log.Printf("Update Blog invoked with: %v\n", inData)

	oid, err := primitive.ObjectIDFromHex(inData.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{
		AuthorID: inData.AuthorId,
		Title:    inData.Title,
		Content:  inData.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog item to update",
		)
	}

	return &emptypb.Empty{}, nil

}

func (s *Server) DeleteBlog(ctx context.Context, inData *pb.BlogId) (*empty.Empty, error) {
	log.Printf("Delete blog was called for: %v\n", inData)

	oid, err := primitive.ObjectIDFromHex(inData.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	filter := bson.M{"_id": oid}

	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown error while deleting: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"The Blog ID could not be found",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) ListBlogs(inData *empty.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("List blogs was invoked")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding message: %v\n", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Internal Error: %v\n", err),
		)
	}

	return nil
}
