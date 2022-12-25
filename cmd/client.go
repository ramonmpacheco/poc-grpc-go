package main

import (
	"context"
	"log"

	"github.com/ramonmpacheco/poc-grpc-go/stub"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	client := stub.NewMathServiceClient(connection)
	defer connection.Close()

	request := &stub.SumRequest{
		Sum: &stub.Sum{
			A: 3,
			B: 4,
		},
	}
	res, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("error during request: %v", err)
	}
	log.Println(res)
}
