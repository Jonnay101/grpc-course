package main

import (
	"context"
	"fmt"
	"io"
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

	// doUnarySum(c, 12, 11)
	doPrimeNumberDecomposition(c, 120)
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

func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient, num int32) {
	req := &calculatorpb.PrimeNumDecRequest{
		Num: num,
	}

	res, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("prime number decomposition failed: %v", err)
	}

	for {
		num, err := res.Recv()
		if err == io.EOF {
			break
		}

		fmt.Printf("%d, ", num.GetResult())
	}
}
