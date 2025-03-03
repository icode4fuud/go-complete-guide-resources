# Go - The Complete Guide

This repository contains course resources (code snapshots, slides) for my [Go - The Complete Guide course](https://acad.link/golang).

## How To Use

The [/code](/code/) folder contains multiple code snapshots for every course sections. You can use the code stored in that folder to compare your code to mine in case you're getting stuck when following along with the course.

The [/attachments](/attachments/) folder contains lecture-specific attachments (i.e., files which are directly attached to individual course lectures) like starting project snapshots.

The [/other](/other/) folder contains other resources - e.g, the course slides.

# Golang for gRPC Communication Pattern
protoc --go_out=. --go-grpc_out=. greeter.proto

Clipped from: https://chat.deepseek.com/a/chat/s/56655ea0-54ef-42f2-b10f-8234599e864c
Q: Give me an example of a Golang route for gRPC communication pattern
Certainly! Below is an example of a simple gRPC server and client in Go, demonstrating a unary RPC communication pattern.
1. Define the Protocol Buffers (.proto file)
First, you need to define the service and messages in a .proto file.
proto
Copy
syntax = "proto3";

package example;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
2. Generate Go code from the .proto file
Use the protoc compiler to generate Go code from the .proto file.
bash
Copy
protoc --go_out=. --go-grpc_out=. greeter.proto

This will generate greeter.pb.go and greeter_grpc.pb.go files.
3. Implement the gRPC Server in Go
go
Copy
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "path/to/your/generated/proto" // Update with the correct import path
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    log.Printf("Received: %v", in.GetName())
    return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
4. Implement the gRPC Client in Go
go
Copy
package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "path/to/your/generated/proto" // Update with the correct import path
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}
5. Run the Server and Client
    1. Start the server:
bash
Copy
go run server.go
    1. Run the client:
bash
Copy
go run client.go
Expected Output
Server Output:
Copy
2023/10/10 12:00:00 server listening at [::]:50051
2023/10/10 12:00:05 Received: World
Client Output:
Copy
2023/10/10 12:00:05 Greeting: Hello World
Explanation
    • Server: The server listens on port 50051 and implements the SayHello method defined in the Greeter service.
    • Client: The client connects to the server, sends a HelloRequest with the name "World", and receives a HelloResponse with the message "Hello World".
This is a basic example of a unary RPC in gRPC using Go. You can extend this pattern to include more complex communication patterns like server streaming, client streaming, or bidirectional streaming.
