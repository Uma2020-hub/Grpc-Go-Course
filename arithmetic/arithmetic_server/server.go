package main

import (
	"context"
	"fmt"
	"grpc-go-course/arithmetic/arithmeticpb/arithmetic/arithmeticpb"
	"io"
	"log"
	"math"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Add(ctx context.Context, req *arithmeticpb.ArithmeticRequest) (*arithmeticpb.ArithmeticResponse, error) {
	fmt.Printf("values sent from client : %v", req)
	op1 := req.GetArithmetic().FirstNo
	op2 := req.GetArithmetic().SecondNo
	op3 := req.Arithmetic.FirstNo
	op4 := req.Arithmetic.SecondNo
	fmt.Println(op1, op2, op3, op4)

	r := op3 + op4
	fmt.Println(r)
	result := &arithmeticpb.ArithmeticResponse{
		Result: r,
	}

	return result, nil
}

func (*server) SquareRoot(ctx context.Context, req *arithmeticpb.SquareRootRequest) (*arithmeticpb.SquareRootResponse, error) {
	fmt.Println("Inside Square Root function of server")
	n := req.GetNumber()
	if n < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Got negative number from client %v", n))
		//error part
	} else {
		//sqrt part
		return &arithmeticpb.SquareRootResponse{
			NumberRoot: math.Sqrt(float64(n)),
		}, nil

	}

}

func (*server) Average(stream arithmeticpb.ArithmeticService_AverageServer) error {
	fmt.Println("Inside server Average Function")
	var res int32
	i := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {

			return stream.SendAndClose(&arithmeticpb.AverageResponse{
				Result: float32(res / int32(i)),
			})
		}
		if err != nil {
			log.Fatalf("error in receiving average Request ")
			return err
		}

		res += req.Arithmetic.FirstNo
		i = i + 1

	}

	//return nil
}

func (*server) Maximum(stream arithmeticpb.ArithmeticService_MaximumServer) error {
	fmt.Println("Inside server Maximum function")
	var MaxNo int32 = 0
	var temp int32
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v", err)
		}
		temp = req.GetArithmetic().GetFirstNo()
		fmt.Printf("Received inputs from stream %v \n", temp)
		if temp > MaxNo {
			MaxNo = temp
			sendErr := stream.Send(&arithmeticpb.MaximumResponse{
				Result: strconv.Itoa(int(MaxNo)),
			})
			if sendErr != nil {
				log.Fatalf("Error in maximum server %v", sendErr)
				return sendErr
			}
		}
		// else {
		// 	maxRes = MaxNo
		// }

	}

}

//Add(context.Context, *ArithmeticRequest) (*ArithmeticResponse, error)
func main() {
	fmt.Println("Inside the server main")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen to server %v", err)
	}

	s := grpc.NewServer()
	arithmeticpb.RegisterArithmeticServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v ", err)
	}
}
