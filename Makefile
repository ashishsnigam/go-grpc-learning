gen-proto:
	rm -rf ./proto/*.go
	protoc --go_out=. --go-grpc_out=. ./calculatorpb/calculator.proto 