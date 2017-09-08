package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alinz/xerxes"
	pb "github.com/alinz/xerxes/example/proto"
)

func main() {
	options := []xerxes.Option{
		xerxes.OptCertificateAuthority("../cert/intermediateCA.crt"),
	}

	x, err := xerxes.New(options...)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := x.Grpc().Dial("server", ":10000")
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
