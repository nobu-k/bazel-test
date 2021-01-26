package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/nobu-k/bazel-go-test/goproto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":3456", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)
	msg := "Hello!"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &pb.Message{Message: msg})
	if err != nil {
		log.Fatalf("failed to echo: %v", err)
	}
	log.Printf("Echo: %v", r.GetMessage())
}
