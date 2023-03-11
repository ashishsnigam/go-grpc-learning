package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"
	"io"
	"log"
)

func doPrimes(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("do primes was invoked")

	req := &calculatorpb.PrimesRequest{
		Number: 123456899,
	}
	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("primesa are %d", res.Result)
	}
}
