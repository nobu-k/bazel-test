package main

import (
	"context"
	"log"
	"net"

	pb "github.com/nobu-k/bazel-go-test/goproto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) Echo(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	log.Printf("received message: %s", in.GetMessage())
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3456")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
