package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ashishsnigam/calc/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("calculator client")
	cc, err := grpc.Dial("localhost:7002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("client error is %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalcularorClient(cc)
	params := &calculatorpb.Params{
		A: 10,
		B: 5,
	}
	res, err := c.Add(context.Background(), params)
	if err != nil {
		log.Fatalf("error in add servicce %v", err)
	}
	fmt.Printf("sum is %d", res.Result)

	res, err = c.Sub(context.Background(), params)
	if err != nil {
		log.Fatalf("error in sub servicce %v", err)
	}
	fmt.Printf("diff is %d", res.Result)
}
