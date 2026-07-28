package StarterGin

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"go-spring.org/stdlib/errutil"
)

// TLSConfig is the nested, off-by-default TLS block shared by every starter.
// It is the union of the fields the starters were each declaring on their own:
// a client/server key pair (CertFile/KeyFile), a CA bundle to verify the peer
// (CAFile), the expected peer name (ServerName), and an escape hatch for local
// testing (InsecureSkipVerify).
//
// Embed it under a `tls` key so the bound properties read as e.g.
// `${spring.redis.tls.enabled}`, `${spring.redis.tls.cert-file}`, ...
type TLSConfig struct {
	// Enabled turns TLS on. It defaults to false so a starter never negotiates
	// TLS unless the operator asks for it.
	Enabled bool `value:"${enabled:=false}"`

	// CertFile and KeyFile are the PEM key pair this side presents. Leave both
	// empty when no client/peer certificate is required.
	CertFile string `value:"${cert-file:=}"`
	KeyFile  string `value:"${key-file:=}"`

	// CAFile is a PEM bundle of root CAs used to verify the peer certificate.
	// When empty the host's default root set is used.
	CAFile string `value:"${ca-file:=}"`

	// ServerName overrides the name checked against the peer certificate, useful
	// when dialing by IP or through a service-discovery label.
	ServerName string `value:"${server-name:=}"`

	// InsecureSkipVerify disables peer certificate verification. Intended for
	// local testing only - never enable it in production.
	InsecureSkipVerify bool `value:"${insecure-skip-verify:=false}"`
}

// Build turns the config into a *tls.Config, or (nil, nil) when TLS is
// disabled so a caller can pass the result straight through to a library that
// treats a nil *tls.Config as "no TLS". It loads the key pair and CA bundle
// from disk when provided.
//
// Errors are wrapped with a generic "tls:" prefix; a starter that wants a
// component-specific prefix should wrap the returned error with
// errutil.Explain(err, "redis: ...").
func (c TLSConfig) Build() (*tls.Config, error) {
	if !c.Enabled {
		return nil, nil
	}
	cfg := &tls.Config{
		ServerName:         c.ServerName,
		InsecureSkipVerify: c.InsecureSkipVerify,
	}
	if c.CertFile != "" || c.KeyFile != "" {
		cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)
		if err != nil {
			return nil, errutil.Explain(err, "tls: failed to load key pair")
		}
		cfg.Certificates = []tls.Certificate{cert}
	}
	if c.CAFile != "" {
		pem, err := os.ReadFile(c.CAFile)
		if err != nil {
			return nil, errutil.Explain(err, "tls: failed to read CA file")
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, errutil.Explain(nil, "tls: no certificates found in CA file %s", c.CAFile)
		}
		cfg.RootCAs = pool
	}
	return cfg, nil
}
