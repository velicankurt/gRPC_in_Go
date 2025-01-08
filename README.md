# gRPC_in_Go

#### protoc -Igreet/proto --go_out=. --go_opt=module={YOUR_MODULE_NAME} --go-grpc_out=. --go-grpc_opt=module={YOUR_MODULE_NAME} greet/proto/dummy.proto
You can Find your module name in go.mod file

With this command you can create go files from proto files