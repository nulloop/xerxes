package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alinz/xerxes"
	pb "github.com/alinz/xerxes/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	security, err := xerxes.NewSecurity("../cert/intermediateCA.crt", "../cert/intermediateCA.pub", "../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := security.ClientTLS("server")

	conn, err := grpc.Dial(":10000", grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
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
