package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jonnay101/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctc context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum called with: %v", req)

	num1 := req.Num1
	num2 := req.Num2

	sum := num1 + num2
	res := &calculatorpb.SumResponse{
		Result: sum,
	}

	return res, nil
}

func main() {
	fmt.Println("Yo, calc server up!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
