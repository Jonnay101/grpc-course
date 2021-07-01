package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jonnay101/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC connection started...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client connection failed: %v", err)
	}
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnarySum(c, 12, 11)
}

func doUnarySum(c calculatorpb.CalculatorServiceClient, num1, num2 int32) {
	req := &calculatorpb.SumRequest{
		Num1: num1,
		Num2: num2,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("sum failed: %v", err)
	}

	fmt.Printf("The sum result is: %d", res.Result)
}
