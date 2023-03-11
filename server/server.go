package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ashishsnigam/calc/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type calcServer struct {
	calculatorpb.UnimplementedCalcularorServer
}

func (*calcServer) Add(ctx context.Context, req *calculatorpb.Params) (*calculatorpb.Result, error) {
	fmt.Printf("Received sum request rpc: %v\n", req)
	firstNumber := req.A
	secondNumber := req.B

	sum := firstNumber + secondNumber
	res := &calculatorpb.Result{Result: sum}
	return res, nil
}

func (*calcServer) Sub(ctx context.Context, req *calculatorpb.Params) (*calculatorpb.Result, error) {
	fmt.Printf("Recieved sub rpc request : %v\n", req)
	if req.A > req.B {
		diff := req.A - req.B
		return &calculatorpb.Result{Result: diff}, nil
	}
	return nil, fmt.Errorf("first number should be big")
}

// Add this calcServer in grpc calcServer
func main() {
	fmt.Println("calculator calcServer")

	lis, err := net.Listen("tcp", "localhost:7002")
	if err != nil {
		log.Fatalf("failed to listen calcServer with error %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalcularorServer(s, &calcServer{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start calcServer with error %v", err)
	}
}
