package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/arielcr/grpc-go-demo/sum/sumpb"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v \n", req)

	first := req.GetNumbers().GetFirst()

	second := req.GetNumbers().GetSecond()

	total := first + second

	res := &sumpb.SumResponse{
		Total: total,
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

	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
