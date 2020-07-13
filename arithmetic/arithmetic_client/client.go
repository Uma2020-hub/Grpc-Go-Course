package main

import (
	"context"
	"fmt"
	"grpc-go-course/arithmetic/arithmeticpb/arithmetic/arithmeticpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("inside client main function")
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server %v", err)
	}
	defer conn.Close()

	c := arithmeticpb.NewArithmeticServiceClient(conn)
	//additionOperation(c)
	//doClientStreaming(c)
	//MaximumOperation(c)
	additionErrorOperation(c)
}

func doClientStreaming(con arithmeticpb.ArithmeticServiceClient) {
	fmt.Println("Inside Client Streaming")
	req := []*arithmeticpb.AverageRequest{
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 1,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 2,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 3,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 4,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 5,
			},
		},
	}
	stream, err := con.Average(context.Background())

	if err != nil {
		log.Fatalf("error in calling average function %v", err)
	}

	for _, r := range req {
		fmt.Printf("Sending number %v", r)
		stream.Send(r)
		time.Sleep(1 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error in receiving client streaming %v", err)
	}
	fmt.Printf("Response from Average Response %v", res.GetResult())
}

//MaximumOperation in client side
func MaximumOperation(con arithmeticpb.ArithmeticServiceClient) {
	fmt.Println("starting to do a bidirectional streaming")
	req := []*arithmeticpb.MaximumRequest{
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 1,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 5,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 3,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 2,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 6,
			},
		},
		{
			Arithmetic: &arithmeticpb.Arithmetic{
				FirstNo: 50,
			},
		},
	}

	waitc := make(chan struct{})
	stream, err := con.Maximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling MAximum  RPC %v", err)
	}
	go func() {

		for _, r := range req {
			stream.Send(r)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	//receive
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error in receiving the result from maximum server %v ", err)
			}
			//maximum := res.Result

			fmt.Printf("Response from maximum server %v \n", res.GetResult())
		}
		close(waitc)
	}()
	// block until everything is done
	<-waitc
}

func additionErrorOperation(con arithmeticpb.ArithmeticServiceClient) {

	fmt.Println("Inside additionErrorOperation Function")
	doErrorCall(con, 12)
	doErrorCall(con, -6)

}

func doErrorCall(con arithmeticpb.ArithmeticServiceClient, num int32) {
	req := &arithmeticpb.SquareRootRequest{
		Number: num,
	}
	res, err := con.SquareRoot(context.Background(), req)

	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			fmt.Println(resErr.Code())
			fmt.Println(resErr.Message())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Printf("We have probably sent a negative number %v \n", num)
			} else {
				fmt.Printf("Something went wrong while calling sqrtfunction %v \n", err)
			}
		}
		//log.Fatalf("error while calling Arithmetic sqrt %v", err)
		return
	}
	log.Printf("Response from Arithmetic Sqrt function : %v %v \n", num, res.GetNumberRoot())
}

func additionOperation(con arithmeticpb.ArithmeticServiceClient) {

	fmt.Println("Inside additionOperation Function")
	req := &arithmeticpb.ArithmeticRequest{
		Arithmetic: &arithmeticpb.Arithmetic{
			FirstNo:  16,
			SecondNo: 98,
		},
	}
	res, err := con.Add(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Arithmetic add %v", err)

	}
	log.Printf("Response from Arithmetic Service : %d", res.Result)

}
