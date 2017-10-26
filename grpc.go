package xerxes

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// DialFn is a function which creates grpc.ClientConn
type DialFn func(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error)

// Grpc is the base for creating grpc client and server
type Grpc struct {
	security *Security
}

// CreateDialFn is the low level, You can cache DialFn and keep using it
func (g *Grpc) CreateDialFn(serverName string) DialFn {
	tlsConfig := g.security.ClientTLS(serverName)

	config := credentials.NewTLS(tlsConfig)

	return func(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
		opts = append([]grpc.DialOption{
			grpc.WithTransportCredentials(config),
		}, opts...)

		return grpc.Dial(target, opts...)
	}
}

// Dial is used to make a grpc client connection with default TLS cert
func (g *Grpc) Dial(serverName, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return g.CreateDialFn(serverName)(target, opts...)
}

// DialWithJwtToken is the same as Dial function except it adds token as well
func (g *Grpc) DialWithJwtToken(token Token, serverName, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append([]grpc.DialOption{
		grpc.WithPerRPCCredentials(token),
	}, opts...)

	return g.CreateDialFn(serverName)(target, opts...)
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
