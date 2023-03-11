package main

import (
	"calc/calculatorpb"
	"fmt"
	"io"
	"log"
)

func (s *calcServer) Max(stream calculatorpb.CalculatorService_MaxServer) error {
	fmt.Println("Bidi streaming example")

	// infinite loop to receive client requests
	lastMax := int32(0)
	for {
		// first collect client inputs
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
		}
		// second - run logic on those inputs
		if req.Number > lastMax {
			lastMax = req.Number
			fmt.Println("max number is ", lastMax)
			// third - send result as response
			result := &calculatorpb.MaxResponse{
				Result: lastMax,
			}
			err = stream.Send(result)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	return nil
}
