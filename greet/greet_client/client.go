package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arielcr/grpc-go-demo/greet/greetpb"

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

	c := greetpb.NewGreetServiceClient(cc)

	// Unary API
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	request := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ariel",
			LastName:  "Orozco",
		},
	}

	response, err := c.Greet(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", response.Result)
}
