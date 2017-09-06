package xerxes

import (
	"crypto/rsa"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

// Token is an interface to describe how we can
// store services token
type Token interface {
	credentials.PerRPCCredentials
	jwt.Claims
}

// token is simple implementation
type token struct {
	value string
}

func (t token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": t.value,
	}, nil
}

// RequireTransportSecurity is being used by higher level code in grpc
func (token) RequireTransportSecurity() bool {
	return true
}

// Valid has to be here for jwt
func (token) Valid() error {
	return nil
}

// NewToken gets an already created and signed token
func NewToken(value string) (Token, error) {
	return &token{value: value}, nil
}

// NewTokenWithSign accepts a map and a private key and create Token and sign it with
// given private key
func NewTokenWithSign(data map[string]interface{}, privateKey *rsa.PrivateKey) (Token, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("need to provide private key to sign the token")
	}

	claims := jwt.MapClaims(data)
	alg := jwt.GetSigningMethod("RS256")
	jwtToken := jwt.NewWithClaims(alg, claims)

	value, err := jwtToken.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &token{value: value}, nil
}
