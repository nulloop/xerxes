package xerxes

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

// Security this common object to be used for GRPC
// it can also be used for other tls related, such as connecting to
// NATs server
type Security struct {
	cp     *x509.CertPool
	cert   *tls.Certificate
	pubKey *rsa.PublicKey
}

// ParseJwt tries to parse and check the sign value
func (s *Security) ParseJwt(value Token) (*jwt.Token, error) {
	tok, ok := value.(*token)
	if !ok {
		return nil, fmt.Errorf("Token is not JwtToken")
	}

	jwtToken, err := jwt.Parse(tok.value, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.pubKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return jwtToken, nil
}

// ClientTLS creates tls config for specific server
// if you want to connect to NATS, make sure the serverName is set to empty
// or you will get `x509: certificate is not valid for any names` error
func (s *Security) ClientTLS(serverName string) *tls.Config {
	tlsConfig := tls.Config{RootCAs: s.cp}
	if serverName != "" {
		tlsConfig.ServerName = serverName
	}

	if s.cert != nil {
		tlsConfig.Certificates = []tls.Certificate{*s.cert}
	}
	return &tlsConfig
}

// ServerTLS returns the certificate to be used mainly in GRPC server
func (s *Security) ServerTLS() *tls.Certificate {
	return s.cert
}

func (s *Security) loadCertificateAuthorityPublicKey(filename string) error {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return fmt.Errorf("public key not encoded properly")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return fmt.Errorf("given public key not valid")
	}

	s.pubKey = pkey
	return nil
}

func (s *Security) loadCertificateAuthority(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	s.cp = x509.NewCertPool()
	if !s.cp.AppendCertsFromPEM(b) {
		return fmt.Errorf("credentials: failed to append certificates")
	}

	return nil
}

func (s *Security) loadCertificate(crt, key string) error {
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return err
	}

	s.cert = &cert
	return nil
}
