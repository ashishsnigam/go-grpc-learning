### grpc protobuf demo code using golang

- Very basic code to help me understand all 4 flows of grpc
   - Unary
    - Server streaming
    - Client streaming
    - Bi-Directional streaming
    
- To run this code
    - install go, protoc compiler and go-grpc bindings
    - run `make` command and code should compile
    - `cd server` folder and run server binary using `./server` and keep it running
    - `cd client` folder and run client binary to test end to end flow
    
- For simplicity, a single `proto` file is used and all rpc calls and messages are put there only