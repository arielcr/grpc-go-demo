package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arielcr/grpc-go-demo/sum/sumpb"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello! I'm the Client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// this will be executed at the end of the code
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)

	// Unary API
	doUnary(c)

}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	request := &sumpb.SumRequest{
		Numbers: &sumpb.Numbers{
			First:  10,
			Second: 20,
		},
	}

	response, err := c.Sum(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", response.Total)
}
