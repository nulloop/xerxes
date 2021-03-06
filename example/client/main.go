package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nulloop/xerxes"
	pb "github.com/nulloop/xerxes/example/proto"
)

func main() {
	options := []xerxes.Option{
		xerxes.OptCertificateAuthority("../cert/intermediateCA.crt"),
		xerxes.OptCertificate("../cert/client.crt", "../cert/client.key"),
	}

	x, err := xerxes.New(options...)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := x.Grpc.Dial("server", ":10000")
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
