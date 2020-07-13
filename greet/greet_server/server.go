package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	//"google.golang.org/grpc/internal/status"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}
func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	for i := 0; i <= 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client cancelled the request !")
			return nil, status.Error(codes.Canceled, "Client cancelled the request")
		}
		time.Sleep(1000 * time.Second)

	}
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.Greeting.FirstName
	//GetGreeting().GetFirstName()
	for i := 0; i < 12; i++ {
		result := "Hello " + firstName + " this is count " + strconv.Itoa(i)
		resStream := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(resStream)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Long Greet was invoked with a streaming Request")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error in long greet %v ", err)
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + " ! "

	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println(" Greet everyone was invoked with a streaming Request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client Stream %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&greetpb.GreetEveryoneReponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client Stream %v", sendErr)
			return sendErr
		}

	}
}

func main() {
	fmt.Println("Testing Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	//register service using reflection
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}

}
