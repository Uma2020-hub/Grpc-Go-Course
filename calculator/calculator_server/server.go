package main

import (
	"fmt"
	"grpc-go-course/calculator/calculatorpb/calculator/calculatorpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) PrimeComposition(req *calculatorpb.PrimeCompositionRequest, stream calculatorpb.PrimeCompositionService_PrimeCompositionServer) error {
	fmt.Println("Inside the server streaming function")
	firstno := req.Primecompose.FirstNo
	var k int32 = 2
	//var factor int32 = 0
	for firstno > 1 {

		if firstno%k == 0 {
			//factor = k
			//res := strconv.Itoa(int(k)) + ","
			resStream := &calculatorpb.PrimeCompositionResponse{
				Result: k,
			}

			stream.Send(resStream)
			time.Sleep(1000 * time.Millisecond)

			firstno = firstno / k

		} else {
			k = k + 1
			fmt.Printf("Divisor incremented by 1 %v", k)
		}

	}
	return nil
}

func main() {
	fmt.Println("Inside the main function in server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error in connecting the server %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterPrimeCompositionServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to register server %v", err)
	}

}
