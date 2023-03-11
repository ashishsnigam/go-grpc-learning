package main

import (
	"calc/calculatorpb"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

func doMax(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("doMax was invoked")
	// call rpc method
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// since it is bidirectional, so we need to have two go routines to handle requests from client and responses from server
	waitc := make(chan struct{})

	// first go-routine to send client data to server
	go func() {
		// prepare requests to be sent from client to server
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("sending no %v\n", number)
			stream.Send(&calculatorpb.MaxRequest{Number: number})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// second go-routine to receive response from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				break
			}
			fmt.Println("Received a new maximum ", res.Result)
		}

		close(waitc)
	}()

	<-waitc // main goroutine to wait till close(waitc) is called
}
