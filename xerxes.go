package xerxes

import jwt "github.com/dgrijalva/jwt-go"

// Option helps chnaging Options
type Option func(*Xerxes) error

// OptCertificate you need this if you want to create grpc server
func OptCertificate(certPath, keyPath string) Option {
	return func(xerxes *Xerxes) error {
		return xerxes.Security.loadCertificate(certPath, keyPath)
	}
}

// OptCertificateAuthorityPublicKey you need this if you want to use token
func OptCertificateAuthorityPublicKey(certPubPath string) Option {
	return func(xerxes *Xerxes) error {
		return xerxes.Security.loadCertificateAuthorityPublicKey(certPubPath)
	}
}

// OptCertificateAuthority you need this if you want to create grpc client
func OptCertificateAuthority(caPath string) Option {
	return func(xerxes *Xerxes) error {
		return xerxes.Security.loadCertificateAuthority(caPath)
	}
}

// Xerxes is a base object
type Xerxes struct {
	Security *Security
	Grpc     *Grpc
}

// Token parse jwt token based on public key
func (x *Xerxes) Token(token Token) (*jwt.Token, error) {
	return x.Security.ParseJwt(token)
}

// New creates xerxes object based on list of options
func New(options ...Option) (*Xerxes, error) {
	security := Security{}

	xerxes := Xerxes{
		Security: &security,
		Grpc: &Grpc{
			security: &security,
		},
	}

	for _, option := range options {
		if err := option(&xerxes); err != nil {
			return nil, err
		}
	}

	return &xerxes, nil
}
