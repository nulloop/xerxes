package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/alinz/xerxes/example/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, user *pb.User) (*pb.Response, error) {
	return &pb.Response{
		Message: fmt.Sprintf("hello %s", user.Name),
	}, nil
}

func main() {
	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("Failed to setup tls: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
	)

	server := server{}

	pb.RegisterGreetingServer(grpcServer, &server)

	// start listening to the network
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = grpcServer.Serve(ln)
	if err != nil {
		panic(err)
	}
}
