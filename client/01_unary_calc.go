package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"
	"log"
)

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("do unary was invoked")

	params := &calculatorpb.AddRequest{
		FirstParam:  10,
		SecondParam: 5,
	}
	res, err := c.Add(context.Background(), params)
	if err != nil {
		log.Fatalf("error in add rpc call %v", err)
	}
	fmt.Printf("\nsum is %d", res.Result)
}
