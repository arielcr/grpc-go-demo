package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doUnary(c)

	// Streaming Server API
	// doServerStreaming(c)

	// Streaming Client API
	doClientStreaming(c)

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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	request := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ariel",
			LastName:  "Orozco",
		},
	}

	stream, err := c.GreetManyTimes(context.Background(), request)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			// We've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "George",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Paul",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "James",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sam",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving repsonse: %v", err)
	}

	fmt.Printf("LongGreet Response: %v\n", response)

}
