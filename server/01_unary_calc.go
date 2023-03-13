package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*calcServer) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Printf("Received sum request rpc: %v\n", req)
	firstNumber := req.FirstParam
	secondNumber := req.SecondParam

	if firstNumber < 0 || secondNumber < 0 {
		return nil, status.Error(codes.InvalidArgument, "both numbers should be positive integers")
	}
	sum := firstNumber + secondNumber
	res := &calculatorpb.AddResponse{Result: sum}
	return res, nil
}
