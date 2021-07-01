package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnay101/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("yo.... server here!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 50051 - default grpc port
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
