package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)
import pb "api/proto"

type SearchService struct {
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println("init success!", time.Now())

	server.Serve(lis)
	log.Println("init success!", time.Now())
}
