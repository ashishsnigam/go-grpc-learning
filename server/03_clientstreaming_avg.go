package main

import (
	"calc/calculatorpb"
	"fmt"
	"io"
	"log"
)

func (s *calcServer) Avg(stream calculatorpb.CalculatorService_AvgServer) error {
	fmt.Println("Avg function was invoked")
	cnt := int64(0)
	sum := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := sum / cnt
			return stream.SendAndClose(&calculatorpb.AvgResponse{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("recieved number ", req.Number)

		sum += req.Number
		cnt++
	}
}
