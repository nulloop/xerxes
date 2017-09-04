package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/alinz/xerxes/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../cert/IntermediateCA.crt", "server")
	if err != nil {
		log.Fatalf("cert load error: %s", err)
	}

	conn, err := grpc.Dial(":10000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetingClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.User{Name: "Ali"})
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Println(resp.Message)
}
