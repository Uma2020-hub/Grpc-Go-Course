package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb/calculator/calculatorpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Inside calculator  client ")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to client %v", err)
	}
	c := calculatorpb.NewPrimeCompositionServiceClient(conn)
	doServerStreaming(c)
}

func doServerStreaming(con calculatorpb.PrimeCompositionServiceClient) {
	fmt.Println("Inside do client streaming")
	//var i int32 = 120
	req := &calculatorpb.PrimeCompositionRequest{
		Primecompose: &calculatorpb.PrimeCompose{
			FirstNo:  210,
			SecondNo: 1,
		},
	}
	resStream, err := con.PrimeComposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not connect to Prime Composition")
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving service call %v", err)
		}
		log.Printf("Response from Prime Composition service %v", msg.GetResult())
	}

}
