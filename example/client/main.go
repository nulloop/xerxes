package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alinz/xerxes"
	pb "github.com/alinz/xerxes/example/proto"
)

func main() {
	security, err := xerxes.NewSecurity("../cert/intermediateCA.crt", "../cert/intermediateCA.pub", "../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	grpc := xerxes.NewGrpc(security)

	conn, err := grpc.Dial("server", ":10000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewGreetingClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.User{Name: "Ali"})
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Println(resp.Message)
}
