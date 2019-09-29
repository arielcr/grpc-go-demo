package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/arielcr/grpc-go-demo/greet/greetpb"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Sum function was invoked with %v: \n", req)

	firstname := req.GetGreeting().GetFirstName()

	result := "Hello " + firstname

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {

	fmt.Println("Server listening...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") // 50051 is gRPC default port

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
