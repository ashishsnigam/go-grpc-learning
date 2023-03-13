package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doAdd(c calculatorpb.CalculatorServiceClient, firstNumber, secondNumber int32) {
	fmt.Println("doAdd was invoked")

	params := &calculatorpb.AddRequest{
		FirstParam:  firstNumber,
		SecondParam: secondNumber,
	}
	res, err := c.Add(context.Background(), params)
	if err != nil {
		// get what kind of grpc error is this
		e, ok := status.FromError(err)
		if ok {
			fmt.Printf("found grpc error %+v", e)
			//fmt.Println(e.Message(), e.Code())
			if e.Code() == codes.InvalidArgument {
				fmt.Printf("\ninvalid parametr is passed in rpc call")
				return
			} else {
				log.Fatalf("some other error in rpc params %v", err)
			}
		}
		log.Fatalf("error in add rpc call %v", err)
	}
	fmt.Printf("\nsum is %d", res.Result)
}
