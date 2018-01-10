package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nulloop/xerxes"
	pb "github.com/nulloop/xerxes/example/proto"
	context "golang.org/x/net/context"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, user *pb.User) (*pb.Response, error) {
	return &pb.Response{
		Message: fmt.Sprintf("hello %s", user.Name),
	}, nil
}

func main() {
	options := []xerxes.Option{
		xerxes.OptCertificate("../cert/server.crt", "../cert/server.key"),
	}

	x, err := xerxes.New(options...)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := x.Grpc.Server()

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
