package xerxes

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Grpc struct {
	security *Security
}

func (g *Grpc) Dial(serverName, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	tlsConfig := g.security.ClientTLS(serverName)

	opts = append([]grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}, opts...)

	return grpc.Dial(target, opts...)
}

func (g *Grpc) Server(opts ...grpc.ServerOption) *grpc.Server {
	cert := g.security.ServerTLS()

	opts = append([]grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(cert)),
	}, opts...)

	return grpc.NewServer(opts...)
}

func NewGrpc(security *Security) *Grpc {
	grpc := Grpc{
		security: security,
	}

	return &grpc
}
