package main

import (
	"calc/calculatorpb"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("calculator client")
	// grpc.Insecure.. is deprecated
	cc, err := grpc.Dial("localhost:7002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("client error is %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	doAdd(c, -1, 10)
	//doPrimes(c)
	//doAvg(c)
	//doMax(c)
}
