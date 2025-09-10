# 30. gRPC in Go

gRPC is a modern, open-source, high-performance RPC (Remote Procedure Call) framework that can run in any environment. It was developed by Google and is built on top of Protocol Buffers and HTTP/2.

## What is gRPC?

gRPC allows a client application to directly call methods on a server application on a different machine as if it were a local object, making it easier to create distributed applications and services.

Key features of gRPC include:
*   **High performance:** gRPC is much more efficient than REST+JSON due to its use of Protobuf and HTTP/2.
*   **Streaming:** gRPC supports streaming, allowing for more flexible and efficient communication patterns.
*   **Language-agnostic:** Like Protobuf, gRPC is language-agnostic.
*   **Strongly typed:** Services and messages are defined in a `.proto` file, providing a clear contract.

## Defining a gRPC Service

You define a gRPC service in a `.proto` file, just like you define messages.

```proto
syntax = "proto3";

package main;

option go_package = "./";

message Person {
  string name = 1;
  int32 id = 2;
}

message GetPersonRequest {
  int32 id = 1;
}

// The Greeter service definition.
service Greeter {
  // Sends a greeting
  rpc GetPerson (GetPersonRequest) returns (Person);
}
```
In this example, we've defined a `Greeter` service with a single RPC method, `GetPerson`. This method takes a `GetPersonRequest` message and returns a `Person` message.

## Generating gRPC Code

After defining your service, you use the `protoc` compiler to generate the gRPC client and server code, just as you did for Protobuf messages.

## Implementing a gRPC Server

Here's how you would implement the `Greeter` service in Go:

```go
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
)

// server is used to implement the Greeter service.
type server struct {
    UnimplementedGreeterServer
}

// GetPerson implements the GetPerson RPC method.
func (s *server) GetPerson(ctx context.Context, in *GetPersonRequest) (*Person, error) {
    log.Printf("Received: %v", in.GetId())
    // In a real application, you would fetch the person from a database.
    return &Person{Name: "Alice", Id: in.GetId()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    RegisterGreeterServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```
In this server:
1.  We define a `server` struct that embeds `UnimplementedGreeterServer` for forward compatibility.
2.  We implement the `GetPerson` method, which is our RPC handler.
3.  In `main`, we create a gRPC server, register our service implementation, and start listening for connections.

## Implementing a gRPC Client

Here's how you would create a client to call the `Greeter` service:

```go
package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := NewGreeterClient(conn)

    // Contact the server and print out its response.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.GetPerson(ctx, &GetPersonRequest{Id: 123})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.GetName())
}
```
In the client:
1.  We use `grpc.Dial()` to connect to the gRPC server.
2.  We create a new `GreeterClient` from the connection.
3.  We call the `GetPerson` RPC method on the client, passing in a context and a request message.

gRPC is a powerful framework for building high-performance, scalable microservices. Its use of Protobuf and HTTP/2 makes it a great choice for modern distributed systems.
