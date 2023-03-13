# most basic proto command without many options
calc:
	rm -rf ./calculatorpb/*.go
	protoc --go_out=. --go-grpc_out=. ./calculatorpb/calculator.proto
	go build -o ./server/server ./server/*.go
	go build -o ./client/client ./client/*.go