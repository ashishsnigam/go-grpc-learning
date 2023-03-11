package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"
	"log"
	"time"
)

func doAvg(c calculatorpb.CalculatorServiceClient) {
	// create an array of input values to send
	inputs := []*calculatorpb.AvgRequest{
		{Number: 10},
		{Number: 20},
		{Number: 30},
		{Number: 40},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < len(inputs); i++ {
		stream.Send(inputs[i])
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("avg is ", res.Result)
}
