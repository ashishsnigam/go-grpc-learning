package main

import (
	"fmt"
	"log"
	"net"

	"calc/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// calcServer is required to have UnimplementedCalculatorServiceServer as per latest grpc changes for forward compatibility
// it can be removed from command line while generating protobuf code but
type calcServer struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

// Add this calcServer in grpc server
func main() {
	fmt.Println("calculator service server code starts at port 7002")
	lis, err := net.Listen("tcp", "localhost:7002")
	if err != nil {
		log.Fatalf("failed to listen calcServer with error %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &calcServer{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start calcServer with error %v", err)
	}
}
