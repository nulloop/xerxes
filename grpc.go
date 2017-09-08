package xerxes

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Grpc is the base for creating grpc client and server
type Grpc struct {
	security *Security
}

// Dial is used to make a grpc client connection with default TLS cert
func (g *Grpc) Dial(serverName, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	tlsConfig := g.security.ClientTLS(serverName)

	opts = append([]grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}, opts...)

	return grpc.Dial(target, opts...)
}

// DialWithJwtToken is the same as Dial function except it adds token as well
func (g *Grpc) DialWithJwtToken(token Token, serverName, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	tlsConfig := g.security.ClientTLS(serverName)

	opts = append([]grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithPerRPCCredentials(token),
	}, opts...)

	return grpc.Dial(target, opts...)
}

// Server creates grpc server based on given options. by default it uses TLS cert
// but it can given `grpc.UnaryInterceptor` for intercepting all grpc calls and
// using that to read the token
func (g *Grpc) Server(opts ...grpc.ServerOption) *grpc.Server {
	cert := g.security.ServerTLS()

	opts = append([]grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(cert)),
	}, opts...)

	return grpc.NewServer(opts...)
}
