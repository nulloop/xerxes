package xerxes

import (
	"context"

	"google.golang.org/grpc/credentials"
)

// Token is an interface to describe how we can
// store services token
type Token interface {
	credentials.PerRPCCredentials
	SignToken() (string, error)
}

// token is simple implementation
type token struct {
	value string
}

func (t *token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	var err error

	token := t.value
	if token == "" {
		token, err = t.SignToken()
		if err != nil {
			return nil, err
		}
	}

	return map[string]string{
		"authorization": token,
	}, nil
}

// RequireTransportSecurity is being used by higher level code in grpc
func (token) RequireTransportSecurity() bool {
	return true
}

// Valid has to be here for jwt
func (t *token) Valid() error {
	return nil
}

func (t *token) SignToken() (string, error) {
	return "", nil
}

func NewJwtToken() {

}
