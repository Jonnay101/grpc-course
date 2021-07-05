package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jonnay101/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello " + firstName + " " + lastName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function invoked with: %v\n", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Hello %s %s, you are number %d", firstName, lastName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
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
