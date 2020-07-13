package main

import (
	"context"
	"fmt"
	"grpc-go-course-master/greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()
	con := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("created Client %f", con)
	//doUnary(con)
	//doServerStreaming(con)
	//doClientStreaming(con)
	//doBiDirectionalStreaming(con)
	doUnaryWithDeadline(con, 5*time.Second)
	doUnaryWithDeadline(con, 1*time.Second)

}

func doBiDirectionalStreaming(con greetpb.GreetServiceClient) {
	fmt.Println("starting to do a client streaming")
	req := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Harini",
				LastName:  "Selva",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "oviya",
				LastName:  "Selva",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Uma",
				LastName:  "Selva",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Selva",
				LastName:  "putthiran",
			},
		},
	}
	waitc := make(chan struct{})
	stream, err := con.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating the stream %v", err)
	}
	//we send a bunch of request to server (go routine)===================
	go func() {
		//function to send bunch of messages
		for _, r := range req {
			fmt.Printf("Sending Message %v", r)
			stream.Send(r)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//we receive a bunch of messages from server (go routine)
	go func() {
		//function to receive bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving %v", err)
				break
			}
			fmt.Printf("Received from BiDirectional Server %v \n", res.GetResult())
		}
		close(waitc)
	}()
	//block until everything is done
	<-waitc
}

func doClientStreaming(con greetpb.GreetServiceClient) {
	fmt.Println("starting to do a client streaming")
	req := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Oviya",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Uma",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Selva",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Harini",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sai",
			},
		},
	}
	stream, err := con.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error in client streaming %v", err)
	}

	for _, r := range req {
		stream.Send(r)
		time.Sleep(1 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error in receiving client streaming%v", err)
	}
	fmt.Printf("Response from LongGreet Response %v\n", res)
}

func doServerStreaming(con greetpb.GreetServiceClient) {
	fmt.Println("Inside unary rpc")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{

			FirstName: "Ganesh",
			LastName:  "Sai",
		},
	}
	resStream, err := con.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greep RPC : %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while calling Greep RPC : %v", err)
		}
		log.Printf("Response from Greet : %v", msg.GetResult())
	}

	//log.Printf("Response from Greet : %v", msg.GetResult())
}

func doUnary(con greetpb.GreetServiceClient) {
	fmt.Println("Inside unary rpc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{

			FirstName: "Ganesh",
			LastName:  "Sai",
		},
	}
	res, err := con.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greep RPC : %v", err)
	}
	log.Printf("Response from Greet : %v", res.Result)
}

func doUnaryWithDeadline(con greetpb.GreetServiceClient, timeOut time.Duration) {
	fmt.Println("Inside doUnaryWithDeadline rpc")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{

			FirstName: "Ganesh",
			LastName:  "Sai",
		},
	}
	//timeout.Seconds()
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	res, err := con.GreetWithDeadline(ctx, req)
	if err != nil {

		resErr, ok := status.FromError(err)

		if ok {
			if resErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline Exceeded")
			} else {
				fmt.Printf("unexpected error %v", resErr)
			}
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC : %v", err)
		}

	}
	log.Printf("Response from GreetWithDeadline : %v", res.Result)
}
