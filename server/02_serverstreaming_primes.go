package main

import (
	"calc/calculatorpb"
	"fmt"
)

func (s *calcServer) Primes(req *calculatorpb.PrimesRequest, stream calculatorpb.CalculatorService_PrimesServer) error {
	fmt.Printf("prime function invoked with %v\n", req)
	number := req.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			res := &calculatorpb.PrimeResponse{
				Result: divisor,
			}
			stream.Send(res)
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
