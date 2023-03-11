package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"
)

func (*calcServer) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Printf("Received sum request rpc: %v\n", req)
	firstNumber := req.FirstParam
	secondNumber := req.SecondParam

	sum := firstNumber + secondNumber
	res := &calculatorpb.AddResponse{Result: sum}
	return res, nil
}
